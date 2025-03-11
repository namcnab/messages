package server

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/google/uuid"
	chatpb "github.com/namcnab/messages/chat"
	emailpb "github.com/namcnab/messages/emailmessages"
)

type EmailMessageServer struct {
	emailpb.UnimplementedEmailMessageServiceServer
}

type ChatMessageServer struct {
	chatpb.UnimplementedChatMessageServiceServer
}

func (s *EmailMessageServer) SendEmail(ctx context.Context, req *emailpb.EmailReq) (*emailpb.GenericResponse, error) {
	//generate uuid
	id := uuid.New().String()

	return &emailpb.GenericResponse{MessageId: []byte(id), Message: req.GetSubject() + " was sent successfully!"}, nil
}

func (s *EmailMessageServer) BroadcastLetter(letter *emailpb.LetterMessage, stream emailpb.EmailMessageService_BroadcastLetterServer) error {
	//generate uuid

	recipients := []string{"user1@example.com", "user2@example.com", "user3@example.com"}

	for _, recipient := range recipients {
		id := uuid.New().String()
		broadcast := &emailpb.Broadcast{
			BroadcastId: []byte(id),
			Sender:      letter.Sender,
			Subscriber:  recipient,
			Message:     letter.Message,
		}
		stream.Send(broadcast)
	}

	return nil

}

func (s *EmailMessageServer) UpdateScribers(stream emailpb.EmailMessageService_UpdateScribersServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed by client")
			return nil
		}
		if err != nil {
			log.Println("Error receiving message:", err)
			return err
		}

		subscribers := req.GetSubscribers()

		for _, subscriber := range subscribers {
			// Process the received message
			log.Println("Received message:", subscriber)
		}

	}
}

func (s *ChatMessageServer) Chat(stream chatpb.ChatMessageService_ChatServer) error {
	//generate uuid
	id := uuid.New().String()
	responseMessage := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed by client")
			return nil
		}
		if err != nil {
			log.Println("Error receiving message:", err)
			return err
		}

		// Process the received message
		log.Println("Received message:", req.GetMessage())

		if strings.Contains(strings.ToLower(req.GetMessage()), "hi") {
			responseMessage = "Hello! How can I help you?"
		} else if strings.Contains(req.GetMessage(), "color") {
			responseMessage = "Green."
		}

		// Send a response
		err = stream.Send(&chatpb.ServerResponse{MessageId: []byte(id), Message: responseMessage})
		if err != nil {
			log.Println("Error sending message:", err)
			return err
		}
	}
}
