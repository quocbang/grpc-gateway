package repositories

import "context"

type Repositories interface {
	Account() Account
	Message() Messaging
}

type Account interface {
	Login(context.Context, LoginRequest) (LoginReply, error)
}

type Messaging interface {
}
