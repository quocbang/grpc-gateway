package connection

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/quocbang/grpc-gateway/server/sender"
)

type EmailSender struct {
	email      *email.Email
	senderInfo EmailSenderConfig
}

type EmailSenderConfig struct {
	SmtpServer  string
	SmtpPort    int
	SenderEmail string
	Password    string
}

func NewEmailSender(esc EmailSenderConfig) (sender.Sender, error) {
	// set up email
	email, err := SetUpEmail(esc)
	if err != nil {
		return nil, fmt.Errorf("failed to set up email, error: %v", err)
	}

	return EmailSender{
		email:      email,
		senderInfo: esc,
	}, nil
}

func SetUpEmail(cfs EmailSenderConfig) (*email.Email, error) {
	return &email.Email{
		From: cfs.SenderEmail,
	}, nil
}

func (es EmailSender) SendEmail() error {
	return es.email.Send(fmt.Sprintf("%s:%d", es.senderInfo.SmtpServer, es.senderInfo.SmtpPort), smtp.PlainAuth("", es.senderInfo.SenderEmail, es.senderInfo.Password, es.senderInfo.SmtpServer))
}
