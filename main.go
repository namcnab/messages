package main

import (
	"log"
	"net"

	chatpb "github.com/namcnab/messages/chat"           // import the generated protobuf code
	emailpb "github.com/namcnab/messages/emailmessages" // import the generated protobuf code
	"github.com/namcnab/messages/server"
	"google.golang.org/grpc"
)

func main() {
	// setup a listener on port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Listening on port 9001")
	}

	// create a new grpc server
	grpcServer := grpc.NewServer()

	log.Println("Registering server...")

	// register our server struct as a handle for the EmailMessageService and ChatMessageService rpc calls that come in through grpcServer
	emailpb.RegisterEmailMessageServiceServer(grpcServer, &server.EmailMessageServer{})
	chatpb.RegisterChatMessageServiceServer(grpcServer, &server.ChatMessageServer{})

	log.Println("Server registered successfully")

	// Serve traffic
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Println("Server started on port 9001")
	}
}
