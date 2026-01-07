package sender

import (
	"context"
	"email-service/internal/model"

	"github.com/wneessen/go-mail"
)

type Sender interface {
	Send(ctx context.Context, email model.Email) error
}

type SmtpSender struct {
	host     string
	port     int
	user     string
	pass     string
	fromAddr string
}

func NewSmtpSender(host string, port int, user, pass, from string) *SmtpSender {
	return &SmtpSender{
		host:     host,
		port:     port,
		user:     user,
		pass:     pass,
		fromAddr: from,
	}
}

func (s *SmtpSender) Send(ctx context.Context, email model.Email) error {
	m := mail.NewMsg()
	if err := m.From(s.fromAddr); err != nil {
		return err
	}
	if err := m.To(s.fromAddr); err != nil { // receiving inbox
		return err
	}
	m.Subject(email.Subject)
	m.SetBodyString(mail.TypeTextPlain, email.Message)
	m.ReplyTo(email.Email) // visitor's email so you can reply directly

	client, err := mail.NewClient(
		s.host,
		mail.WithPort(s.port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithTLSPolicy(mail.TLSMandatory),
	)
	if err != nil {
		return err
	}
	return client.DialAndSend(m)
}