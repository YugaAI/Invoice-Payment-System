package config

import (
	"os"
	"strconv"
	"time"
)

type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func LoadRedisConfig() *RedisConfig {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil
	}
	return &RedisConfig{
		Addr:         getEnv("REDIS_ADDR", "localhost:6379"),
		Password:     os.Getenv("REDIS_PASSWORD"),
		DB:           db,
		DialTimeout:  getDuration("REDIS_DIAL_TIMEOUT", "5s"),
		ReadTimeout:  getDuration("REDIS_READ_TIMEOUT", "3s"),
		WriteTimeout: getDuration("REDIS_WRITE_TIMEOUT", "3s"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getDuration(key, fallback string) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		v = fallback
	}
	d, _ := time.ParseDuration(v)
	return d
}
