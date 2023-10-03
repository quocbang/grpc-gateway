package account

import (
	"context"
	"errors"

	repositoriesErr "github.com/quocbang/grpc-gateway/server/errors"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"gorm.io/gorm"
)

type service struct {
	pg *gorm.DB
}

func NewAccount(pg *gorm.DB) repositories.Account {
	return service{
		pg: pg,
	}
}

func (s service) Login(ctx context.Context, req repositories.LoginRequest) (repositories.LoginReply, error) {
	user := models.User{}
	err := s.pg.Where("username=?", req.Username).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repositories.LoginReply{}, repositoriesErr.Error{
				Code:    repositoriesErr.Code_ACCOUNT_DOES_NOT_EXIST,
				Details: "account does not exist",
			}
		}
		return repositories.LoginReply{}, err
	}
	return repositories.LoginReply{
		Username:     user.Username,
		HashPassword: user.Password,
		Roles:        user.Role,
	}, nil
}
