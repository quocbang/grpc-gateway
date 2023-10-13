package repositories

import (
	"context"
	"database/sql"
)

type Repositories interface {
	Account() Account
	Transactions
}

type Transactions interface {
	BeginTx(context.Context, ...*sql.TxOptions) (Repositories, error)
	Rollback() error
	Commit() error
	Close() error
}

type Account interface {
	Login(context.Context, LoginRequest) (LoginReply, error)
	CreateAccount(context.Context, CreateAccountRequest) error
	GetAccount(context.Context, GetAccountRequest) (GetAccountReply, error)
	CreateVerifyAccount(context.Context, CreateVerifyAccountRequest) error
	GetVerifyAccount(context.Context, GetVerifyAccountRequest) (GetVerifyAccountReply, error)
}
