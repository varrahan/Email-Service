package service

import (
	"context"
	"email-service/internal/model"
	"email-service/internal/sender"

	"go.uber.org/zap"
)

type EmailService struct {
	sender sender.Sender
	logger *zap.Logger
}

func NewEmailService(sender sender.Sender, logger *zap.Logger) *EmailService {
	return &EmailService{
		sender: sender,
		logger: logger,
	}
}

func (s *EmailService) SendContactEmail(ctx context.Context, email model.Email) error {
	s.logger.Info("receiving and sending email",
		zap.String("sender", email.Email),
		zap.String("subject", email.Subject),
	)

	err := s.sender.Send(ctx,email)
	if err != nil {
		s.logger.Error(
			"Failed to send email",
			zap.String("sender", email.Email),
			zap.Error(err),
		)
	}
	return nil
}