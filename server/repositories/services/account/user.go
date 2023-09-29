package account

import (
	"github.com/quocbang/grpc-gateway/server/repositories"
	"gorm.io/gorm"
)

type service struct {
	db gorm.DB
}

func NewAccount(db gorm.DB) repositories.Account {
	return service{
		db: db,
	}
}
