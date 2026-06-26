package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
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
