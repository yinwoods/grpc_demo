package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/bioothod/grpc_test/grtest"
)

func main() {
	addr := flag.String("remote", "0.0.0.0:12345", "Remote address to connect to")
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("Could not connect to %s: %v\n", *addr, err)
	}
	defer conn.Close()

	log.Printf("Connected to %s\n", *addr)

	client := pb.NewTestServiceClient(conn)

	ping := &pb.Ping{
		Ping: "ping",
		Aux:  "some data",
	}

	pong, err := client.PingRequest(context.Background(), ping)
	if err != nil {
		log.Fatalf("Could not ping the server: %v\n", err)
	}
	fmt.Printf("sends:\n\t%+v\n", ping)
	fmt.Printf("receives:\n\t%+v\n", pong)

	stream, err := client.Stream(context.Background(), ping)
	if err != nil {
		log.Fatalf("Could not create server-side steam")
	}

	var prevTime time.Time
	for {
		pong, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed: %v.Stream(%v) = %v, %v", client, ping, pong, err)
		}

		currTime := time.Now()
		if currTime.After(prevTime.Add(10 * time.Second)) {
			fmt.Printf("receives:\n\t%+v\n", pong)
			prevTime = currTime
		}
	}
}
