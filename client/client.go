package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/fly0c8/TLSGrpcDemo/pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"time"
)

var (
	serverAddr = flag.String("server_addr", "localhost:9111", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCommandServiceClient(conn)

	executeclient, err := client.Execute(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			clientmsg, err := executeclient.Recv()
			if err != nil {
				fmt.Printf("Error is: %v\n", err)
				fmt.Println("connection terminmated, needs reconnect")
				return
			} else {
				log.Println("recv from server: ", clientmsg.Msg)
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		err := executeclient.Send(&pb.ServerMsg{Msg: msg})
		if err != nil {
			log.Fatal("client failed to send msg:",err)
		}
		time.Sleep(time.Second)
	}

}
