package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/perbu/go-intro/grpc/pb"
	"github.com/segmentio/kafka-go"
	"log"
)

type server struct {
}

func main() {
	s := server{}
	run(context.Background(), &s)
}

func run(ctx context.Context, s *server) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "add.req",
		Partition: 0,
	})
	r.SetOffset(kafka.LastOffset)
	w := kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "add.resp",
		BatchSize: 1,
	}

	for {
		reqMessage, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("ERR: ", err.Error())
			break
		}
		addReq := pb.CalcRequest{}
		if err := proto.Unmarshal(reqMessage.Value, &addReq); err != nil {
			log.Fatalf("PB PARSE: %s", err)
		}

		res, err := s.Add(&addReq)
		resBytes, err := proto.Marshal(&res)
		respMessage := kafka.Message{Value: resBytes, Key: reqMessage.Key}
		w.WriteMessages(ctx, respMessage)
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
func (s server) Add(request *pb.CalcRequest) (pb.CalcResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	log.Printf("Add: %f + %f = %f", a, b, result)
	return pb.CalcResponse{Result: result}, nil
}

func (s server) Mul(request pb.CalcRequest) (pb.CalcResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return pb.CalcResponse{Result: result}, nil
}
