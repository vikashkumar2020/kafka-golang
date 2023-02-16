package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main(){

	// consumer setup
	conn,_ := kafka.DialLeader(context.Background(),"tcp", "localhost:9092", "topic_test", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second*8))

	// read single message
	message,_ := conn.ReadMessage(1e6)  
	fmt.Println(string(message.Value))

	// read batch of meassage
	batch := conn.ReadBatch(1e1,1e6)
	bytes := make([]byte,1e3)
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}

		fmt.Println(string(bytes))
	}


}