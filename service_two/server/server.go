package server

import (
	"cesargdd/service_two/blogpb"
	"fmt"
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
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		fmt.Println(err)
	}
	return nc
}
