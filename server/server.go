package main

import (
	"flag"
	"fmt"
	"github.com/fly0c8/TLSGrpcDemo/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)
var (
	port = flag.Int("port",9111, "port on which grpc server should listen")
)
type ServerImpl struct {
	pb.UnimplementedCommandServiceServer
}
func (s* ServerImpl) Execute(stream pb.CommandService_ExecuteServer) error {
	go func() {
		i:=0
		for {
			time.Sleep(time.Second*5)
			i++
			stream.Send(&pb.ClientMsg{Msg: fmt.Sprintf("Server send HI%d", i)})
		}
	}()
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return nil
		}
		if err != nil {
			return err
		}
		stream.Send(&pb.ClientMsg{Msg: fmt.Sprintf("Thank you for message: %s", msg)})
	}
	return nil
}

func main() {
	flag.Parse()
	log.Printf("Starting gRPC server on port %d", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCommandServiceServer(grpcServer, &ServerImpl{})
	grpcServer.Serve(listener)
}
