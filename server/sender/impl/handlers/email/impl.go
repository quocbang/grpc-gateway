package email

import (
	"context"
	"log"

	"github.com/jordan-wright/email"
	"go.uber.org/zap"

	"github.com/quocbang/grpc-gateway/pkg/grpc/interceptors"
	"github.com/quocbang/grpc-gateway/server/sender"
)

type emailService struct {
	email *email.Email
	send  func() error
}

func NewEmailService(email *email.Email, send func() error) sender.Emails {
	return emailService{
		email: email,
		send:  send,
	}
}

func (es emailService) SendVerifyEmail(ctx context.Context, to string, subject string, content string) error {
	// setup send info
	es.email.To = []string{to}
	es.email.HTML = []byte(content)
	es.email.Subject = subject

	// do send email
	// chanErr := make(chan error, 1)
	// go func() {
	// 	err := es.send()
	// 	if err != nil {
	// 		chanErr <- err
	// 	}
	// }()
	// go func(c context.Context, err chan error) {
	// 	printErr(ctx, chanErr)
	// 	close(chanErr)
	// }(ctx, chanErr)

	return es.send()
}

// TODO: resolve later because it's another process and not work for log.
func printErr(ctx context.Context, errs <-chan error) {
	for err := range errs {
		if err != nil {
			interceptors.GetLoggerFormContext(ctx).Error("Error during send email", zap.Error(err))
			return
		}
	}
	log.Println("send email successfully")
}
