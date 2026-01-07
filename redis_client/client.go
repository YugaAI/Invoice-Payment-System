package redis_client

import (
	"context"
	"invoice-payment-system/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	ctx, cencel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cencel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("[WARN] redis_client not available: %v", err)
		log.Println("[WARN] App will run WITHOUT Redis cache")
		return nil
	}

	log.Println("[INFO] redis_client connected")
	return client
}
