package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "10.10.10.43:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	for {
		result, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "pay_req_group",
			Consumer: "pay_req_consumer",
			Streams:  []string{"pay_req_stream", ">"},
			Count:    1,
			Block:    0,
		}).Result()

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("result:%+v\n", result)
		for _, message := range result[0].Messages {
			// 处理支付请求数据
			fmt.Printf("Received payment request message: %v\n", message)
			fmt.Println("message.Values:", message.Values)
		}

		// Acknowledge the message
		fmt.Println("len(result[0].Messages):", len(result[0].Messages))
		if len(result[0].Messages) > 0 {
			_, err = client.XAck(ctx, "pay_req_stream", "pay_req_group", result[0].Messages[0].ID).Result()
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
