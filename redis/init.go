package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	Auth     string `yaml:"auth"`
	Database int32  `yaml:"database"`
}
type Session struct {
	TTL int32 `yaml:"ttl"`
}
type Server struct {
	Host    string   `yaml:"host"`
	Port    int32    `yaml:"port"`
	Session *Session `yaml:"session"`
}

var client *redis.Client

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", "10.10.10.43", "6379"),
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	client = rdb
}

func GetRedisClient(ctx context.Context) *redis.Client {
	return client.WithContext(ctx)
}
