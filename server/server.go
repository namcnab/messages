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
