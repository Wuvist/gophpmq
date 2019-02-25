package main

import (
	"context"
	"log"
	"test/msg"

	"github.com/golang/protobuf/proto"
	kafka "github.com/segmentio/kafka-go"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "msg",
		Balancer: &kafka.LeastBytes{},
	})

	m := &msg.Msg{}
	m.Msg = "welcome"
	m.Version = 2019

	data, err := proto.Marshal(m)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: data,
		},
	)

	if err != nil {
		panic(err)
	}

	w.Close()
}
