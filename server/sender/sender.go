package sender

import "context"

type Sender interface {
	Email() Emails
}

type Emails interface {
	SendVerifyEmail(ctx context.Context, to string, subject string, content string) error
}
