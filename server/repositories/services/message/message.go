package message

import (
	"github.com/quocbang/grpc-gateway/server/repositories"
	"gorm.io/gorm"
)

type service struct {
	pg *gorm.DB
}

func NewMessage(pg *gorm.DB) repositories.Messaging {
	return service{
		pg: pg,
	}
}
