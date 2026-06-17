package core

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

/**
 * Main function to connect to postgresql using pgx
 * library and env connection values; first making
 * db pool and then verifying conection is successful
 */
func DBConnection() (*pgxpool.Pool, error) {
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbPool, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		slog.Error("Error connecting api to pgdb",
			"event", "pgdb.connection_pool",
			"name", dbName,
			"user", dbUser,
			"pass", dbPass,
			"port", dbPort,
			"host", dbHost,
			"error", err,
		)

		return nil, err
	}

	if err = dbPool.Ping(context.Background()); err != nil {
		slog.Error("Error verifying db connection",
			"event", "pgdb.connection_ping",
			"error", err,
		)

		return nil, err
	}

	slog.Info("DB connection successful",
		"event", "pgdb.connection",
	)

	return dbPool, nil
}

/**
 * Function to connect valkey server using
 * official valkey go library with env port information
 */
func CacheConnection() (valkey.Client, error) {
	cacheHost := os.Getenv("CACHE_HOST")
	cachePort := os.Getenv("CACHE_PORT")

	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{fmt.Sprintf("%s:%s", cacheHost, cachePort)},
	})

	if err != nil {
		slog.Error("Error connecting api to valkey",
			"event", "cache.valkey_connection",
			"host", cacheHost,
			"port", cachePort,
			"error", err,
		)

		return nil, err
	}

	slog.Info("Connection to cache storage service complete",
		"event", "cache.connection",
	)

	return client, nil
}
