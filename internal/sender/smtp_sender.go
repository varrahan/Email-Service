package sender

import (
	"context"
	"email-service/internal/model"

	"github.com/wneessen/go-mail"
	"go.uber.org/zap"
)

type Sender interface {
	Send(ctx context.Context, email model.Email) error
}

type SmtpSender struct {
	host     string
	port     int
	user     string
	pass     string
	toAddr   string
	fromAddr string
	logger   *zap.Logger
}

func NewSmtpSender(host string, port int, user, pass, to string, from string, logger *zap.Logger) *SmtpSender {
	return &SmtpSender{
		host:     host,
		port:     port,
		user:     user,
		pass:     pass,
		toAddr:	  to,
		fromAddr: from,
		logger:   logger,
	}
}

func (s *SmtpSender) Send(ctx context.Context, email model.Email) error {
	m := mail.NewMsg()
	if err := m.From(s.fromAddr); err != nil {
		s.logger.Error(
			"Error with setting FROM address", 
			zap.String("sender", email.Email),
			zap.Error(err),
		)
		return err
	}
	if err := m.To(s.toAddr); err != nil { 
		s.logger.Error(
			"Error with setting TO address", 
			zap.String("sender", email.Email),
			zap.Error(err),
		)
		return err
	}
	m.Subject(email.Subject)
	m.SetBodyString(mail.TypeTextPlain, email.Message)

	if err := m.ReplyTo(email.Email); err != nil {
		s.logger.Error(
			"Error with setting ReplyTo address", 
			zap.String("sender", email.Email),
			zap.Error(err),
		)
		return err
	}

	client, err := mail.NewClient(
		s.host,
		mail.WithPort(s.port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithUsername(s.user),
		mail.WithPassword(s.pass),
		
	)
	if err != nil {
		s.logger.Error(
			"Error with creating client instance", 
			zap.String("sender", email.Email),
			zap.Error(err),
		)
		return err
	}
	return client.DialAndSend(m)
}