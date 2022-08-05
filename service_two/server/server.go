package server

import (
	"cesargdd/service_two/blogpb"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func BlogSrv() blogpb.BlogServiceClient {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(os.Getenv("SERVER_ONE_URL"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer cc.Close()
	blogsvc := blogpb.NewBlogServiceClient(cc)

	return blogsvc
}

func NatsSvc() *nats.Conn {
	nc, _ := nats.Connect(nats.DefaultURL)
	return nc
}
