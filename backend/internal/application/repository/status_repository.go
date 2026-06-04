package repository

import (
	"canary-stream/backend/core"
	"canary-stream/backend/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

type statusRepository struct {
	db *pgxpool.Pool
	vk valkey.Client
}

func (repository *statusRepository) SelectAll(ctx context.Context) ([]domain.Status, error) {
	const statusIndex string = "statusIndex"
	const expire int64 = 600
	const expireDuration time.Duration = time.Duration(expire * int64(time.Second))

	resp := repository.vk.Do(ctx, repository.vk.B().Smembers().
		Key(statusIndex).
		Build())

	keys, err := resp.AsStrSlice()

	log.Printf("keys: %s", strings.Join(keys, "-"))

	if err == nil && len(keys) > 0 {
		var statusList []domain.Status
		var cacheHit = true

		for _, slugStr := range keys {
			key := fmt.Sprintf("status:%s", slugStr)

			getResp := repository.vk.Do(
				ctx,
				repository.vk.B().Get().Key(key).Build(),
			)

			statusData, err := getResp.AsBytes()

			log.Printf("getting: %s", string(statusData))

			if err != nil {
				cacheHit = false
				break
			}

			var s domain.Status

			if err := json.Unmarshal(statusData, &s); err != nil {
				cacheHit = false
				break
			}

			statusList = append(statusList, s)
		}

		if cacheHit {
			return statusList, nil
		}
	}

	log.Printf("No cache finded")

	rows, err := repository.db.Query(ctx, core.Queries["STATUS_GET_ALL"])

	if err != nil {
		return nil, fmt.Errorf("failed to query status data: %w", err)
	}

	defer rows.Close()

	var status []domain.Status

	for rows.Next() {
		var row domain.Status

		err = rows.Scan(
			&row.Name,
			&row.Slug,
		)

		if err != nil {
			return nil, fmt.Errorf("Failed to scan genre row: %w", err)
		}

		status = append(status, row)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during row iteration: %w", err)
	}

	log.Printf("saving data on cache")

	for _, s := range status {
		jsonData, err := json.Marshal(s)

		log.Printf("data: %s", string(jsonData))

		if err != nil {
			continue
		}

		keySlug := fmt.Sprintf("status:%s", s.Slug)

		repository.vk.Do(
			ctx,
			repository.vk.B().Set().Key(keySlug).Value(string(jsonData)).Ex(expireDuration).Build(),
		)

		repository.vk.Do(
			ctx,
			repository.vk.B().Sadd().Key(statusIndex).Member(s.Slug).Build(),
		)
	}

	repository.vk.Do(
		ctx,
		repository.vk.B().Expire().Key(statusIndex).Seconds(expire).Build(),
	)

	return status, nil
}

func NewStatusRepository(db *pgxpool.Pool, vk valkey.Client) domain.StatusRepository {
	return &statusRepository{
		db: db,
		vk: vk,
	}
}
