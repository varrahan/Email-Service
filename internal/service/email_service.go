package service

import (
	"context"
	"email-service/internal/model"
	"email-service/internal/sender"
)

type EmailService struct {
	sender sender.Sender
}

func NewEmailService(s sender.Sender) *EmailService {
	return &EmailService{sender: s}
}

func (s *EmailService) SendContactEmail(ctx context.Context, email model.Email) error {
	// Here you could validate, log, etc.
	return s.sender.Send(ctx, email)
}