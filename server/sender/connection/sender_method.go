package connection

import (
	"github.com/quocbang/grpc-gateway/server/sender"
	"github.com/quocbang/grpc-gateway/server/sender/impl/handlers/email"
)

func (s EmailSender) Email() sender.Emails {
	return email.NewEmailService(s.email, s.SendEmail)
}
