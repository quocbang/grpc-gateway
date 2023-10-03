package connection

import (
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/services/account"
	"github.com/quocbang/grpc-gateway/server/repositories/services/message"
)

func (db DB) Account() repositories.Account {
	return account.NewAccount(db.Postgres)
}

func (db DB) Message() repositories.Messaging {
	return message.NewMessage(db.Postgres)
}
