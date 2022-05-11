package pub

import (
	"context"
	"fmt"
	"gopubsub/red"
)

func Pub() {
	redisConn := red.GetRedisConnection()
	defer redisConn.Close()

	ctx := context.Background()

	var message string

	for {
		fmt.Print("Write a message to send: ")
		fmt.Scanln(&message)
		redisConn.Publish(ctx, "go", message)
	}
}
