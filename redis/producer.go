package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

func main() {
	ctx := context.Background()
	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.10.10.43:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	// 使用PING命令检查Redis连接是否成功
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis connection failed:", err)
		return
	}

	fmt.Println("Redis connection successful:", pong)

	// 使用INFO命令获取Redis服务器信息
	info, err := rdb.Info(ctx).Result()
	if err != nil {
		fmt.Println("Failed to get Redis server info:", err)
		return
	}

	// 从Redis服务器信息中提取Redis版本号
	// 解析Redis服务器信息，获取Redis版本号
	version := "unknown"
	for _, line := range strings.Split(info, "\r\n") {
		if strings.HasPrefix(line, "redis_version:") {
			version = strings.TrimPrefix(line, "redis_version:")
			break
		}
	}
	fmt.Println("Redis version:", version)

	//// 创建名为"pay_req_stream"的Redis Stream
	//_, err = rdb.XGroupCreateMkStream(ctx, "pay_req_stream", "pay_req_group", "$").Result()
	//if err != nil {
	//	fmt.Println("Failed to create Redis Stream and consumer group:", err)
	//	return
	//}
	//
	//fmt.Println("Redis Stream and consumer group created successfully.")

	// 使用XRANGE命令获取Redis Stream列表
	streams, err := rdb.XRange(ctx, "pay_req_stream", "-", "+").Result()
	if err != nil {
		fmt.Println("Failed to get Redis Stream list:", err)
		return
	}

	// 输出Redis Stream列表
	fmt.Println("Redis Stream list:")
	for _, stream := range streams {
		fmt.Println("-", stream.ID)
	}

	// 向名为"pay_req_stream"的Redis Stream中添加一条消息
	msgID, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "pay_req_stream",
		ID:     "*",
		Values: map[string]interface{}{
			"message": "Hello, world!",
			"id":      1,
		},
	}).Result()
	if err != nil {
		fmt.Println("Failed to add message to Redis Stream:", err)
		return
	}

	fmt.Println("Message added to Redis Stream successfully with ID:", msgID)

	// 将消息发送到名为"pay_req_group"的消费者组中
	_, err = rdb.XGroupSetID(ctx, "pay_req_stream", "pay_req_group", msgID).Result()
	if err != nil {
		fmt.Println("Failed to send message to consumer group:", err)
		return
	}

	fmt.Println("Message sent to consumer group successfully.")
}
