package cache

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/valkey-io/valkey-go"
)

/**
 * Function to connect valkey server using
 * official valkey go library with env port information
 */
func CacheConnection() (valkey.Client, error) {
	cacheHost := os.Getenv("CACHE_HOST")
	cachePort := os.Getenv("CACHE_PORT")
	cacheUser := os.Getenv("VALKEY_USER")
	cachePass := os.Getenv("VALKEY_PASSWORD")

	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{fmt.Sprintf("%s:%s", cacheHost, cachePort)},
		Username:    cacheUser,
		Password:    cachePass,
	})

	if err != nil {
		slog.Error("Error connecting api to valkey",
			"event", "cache.valkey_connection",
			"host", cacheHost,
			"port", cachePort,
			"user", cacheUser,
			"error", err,
		)

		return nil, err
	}

	slog.Info("Connection to cache storage service complete",
		"event", "cache.connection",
	)

	return client, nil
}
