package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/perbu/go-intro/grpc/pb"
	"github.com/segmentio/kafka-go"
	"log"
	"math/rand"
	"time"
)

func add(ctx context.Context, a, b float32) (res float32) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "add.resp",
		Partition: 0,
	})
	err := r.SetOffset(kafka.LastOffset)
	if err != nil {
		log.Fatalf("set offset: %s")
	}
	w := kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "add.req",
		BatchSize: 1,
	}
	log.Println("Reader and writer ok.")
	req := pb.CalcRequest{
		A: a,
		B: b,
	}
	reqBytes, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("marshal err: %s", err)
	}
	kMsg := kafka.Message{Value: reqBytes}
	err = w.WriteMessages(ctx, kMsg)
	if err != nil {
		log.Fatalf("write err: %s", err)
	}
	log.Println("RPC sent ---> server, waiting for response")
	responseMessage, err := r.ReadMessage(ctx)
	if err != nil {
		log.Fatalf("read err: %s", err)
	}
	log.Println("Got reply!")
	var resp pb.CalcResponse
	err = proto.Unmarshal(responseMessage.Value, &resp)
	if err != nil {
		log.Fatalf("unmarshal err: %s", err)
	}

	return resp.GetResult()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a:= rand.Float32()
	b := rand.Float32()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res := add(ctx, a, b)
	cancel()
	fmt.Printf("RPC: %f + %f = %f\n", a, b, res)
	fmt.Printf("local: %f + %f = %f\n", a, b, a+b)

}
