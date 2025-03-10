package server

import (
	"context"

	"github.com/google/uuid"
	emailpb "github.com/namcnab/messages/emailmessages"
)

type EmailMessageServer struct {
	emailpb.UnimplementedEmailMessageServiceServer
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
