package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	// producer setup
	conn,_ := kafka.DialLeader(context.Background(),"tcp", "localhost:9092", "topic_test", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second*10))

	conn.WriteMessages(kafka.Message{Value:[]byte("hello-kafka from me")}) // write single message at a time
}