package repository

import (
	"canary-stream/backend/core"
	"canary-stream/backend/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

type statusRepository struct {
	db *pgxpool.Pool
	vk valkey.Client
}

func vcStatus(ctx context.Context, repository *statusRepository, statusKey string) ([]domain.Status, error) {
	var statusData []domain.Status
	var cacheHit bool = true

	keys, err := repository.vk.
		Do(ctx, repository.vk.B().Smembers().Key(statusKey).Build()).
		AsStrSlice()

	if err != nil || len(keys) < 1 {
		log.Printf("Error getting status keys: %v, len: %v", err, len(keys))
		return nil, fmt.Errorf("Error getting status index values: %w", err)
	}

	for i, slug := range keys {
		var statusSchema domain.Status

		key := fmt.Sprintf("status:%s", slug)

		value, err := repository.vk.
			Do(ctx, repository.vk.B().Get().Key(key).Build()).
			AsBytes()

		if err != nil {
			log.Printf("Error mapping status values at %v: %v", i, err)
			cacheHit = false
			break
		}

		if err := json.Unmarshal(value, &statusSchema); err != nil {
			cacheHit = false
			break
		}

		statusData = append(statusData, statusSchema)
	}

	if !cacheHit {
		return nil, fmt.Errorf("Error mapping status data")
	}

	return statusData, nil
}

func scStatus(ctx context.Context, repository *statusRepository, status []domain.Status, statusKey string) {
	const expire int64 = 600
	const duration time.Duration = time.Duration(expire * int64(time.Second))

	for i, state := range status {
		data, err := json.Marshal(state)

		if err != nil {
			log.Printf("Error saving on cache row %v, %v: %v", i, state.Slug, err)
			continue
		}

		key := fmt.Sprintf("status:%s", state.Slug)
		value := string(data)

		repository.vk.Do(ctx, repository.vk.B().Set().Key(key).Value(value).Ex(duration).Build())
		repository.vk.Do(ctx, repository.vk.B().Sadd().Key(statusKey).Member(state.Slug).Build())
	}

	repository.vk.Do(ctx, repository.vk.B().Expire().Key(statusKey).Seconds(expire).Build())
}

func (repository *statusRepository) SelectAll(ctx context.Context) ([]domain.Status, error) {
	const statusIndex string = "statusIndex"
	var status []domain.Status
	var rowHit bool = true

	data, err := vcStatus(ctx, repository, statusIndex)

	if err == nil {
		return data, nil
	}

	log.Printf("No status data finded on cache")
	rows, err := repository.db.Query(ctx, core.Queries["STATUS_GET_ALL"])

	if err != nil {
		log.Printf("Failed to query status data: %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var state domain.Status

		err = rows.Scan(
			&state.Name,
			&state.Slug,
		)

		if err != nil {
			log.Printf("Failed to patch status values for row: %v", err)
			rowHit = false
			break
		}

		status = append(status, state)
	}

	if !rowHit || rows.Err() != nil {
		return nil, fmt.Errorf("Error scanning status db data")
	}

	scStatus(ctx, repository, status, statusIndex)

	return status, nil
}

func NewStatusRepository(db *pgxpool.Pool, vk valkey.Client) domain.StatusRepository {
	return &statusRepository{
		db: db,
		vk: vk,
	}
}
