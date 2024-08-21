package repositories

import (
	"context"
	"database/sql"

	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
)

type Repositories interface {
	Account() Account
	Product() Product
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
	UpdateUserRole(context.Context, UpdateUserRoleRequest) (CommonUpdateReply, error)

	CreateVerifyAccount(context.Context, CreateVerifyAccountRequest) error
	GetUnVerifyAccount(context.Context, GetUnVerifyAccountRequest) (GetUnVerifyAccountReply, error)
	UpdateVerifiedAccount(context.Context, UpdateVerifiedAccountRequest) (CommonUpdateReply, error)

	CreateSessions(context.Context, CreateSessionsRequest) error
	GetSessions(context.Context, GetSessionsRequest) (GetSessionsReply, error)
}

type Product interface {
	Create(context.Context, models.Product) error
	Creates(context.Context, []models.Product) error
}
