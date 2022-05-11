package sub

import (
	"context"
	"fmt"
	"gopubsub/red"
)

func Sub() {
	redisConn := red.GetRedisConnection()
	defer redisConn.Close()

	ctx := context.Background()
	subChannel := redisConn.Subscribe(ctx, "go").Channel()

	for message := range subChannel {
		fmt.Printf("Channel %s: %s\n", message.Channel, message.Payload)
	}
}
