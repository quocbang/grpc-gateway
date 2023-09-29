package message

import (
	"github.com/quocbang/grpc-gateway/server/repositories"
	"gorm.io/gorm"
)

type service struct {
	db gorm.DB
}

func NewMessage(db gorm.DB) repositories.Messaging {
	return service{
		db: db,
	}
}
