package account

import (
	"context"
	"errors"

	"github.com/quocbang/grpc-gateway/server/repositories"
	repositoriesErr "github.com/quocbang/grpc-gateway/server/repositories/errors"
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
	user := models.Account{}
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

func (s service) CreateAccount(ctx context.Context, req repositories.CreateAccountRequest) error {
	account := models.Account{
		Username: req.Username,
		Email:    req.Email,
		Password: req.HashPassword,
	}
	return s.pg.Create(&account).Error
}

func (s service) GetAccount(ctx context.Context, req repositories.GetAccountRequest) (repositories.GetAccountReply, error) {
	account := models.Account{}
	if err := s.pg.Where("username = ?", req.Username).Take(&account).Error; err != nil {
		return repositories.GetAccountReply{}, err
	}
	return repositories.GetAccountReply{Account: account}, nil
}

func (s service) CreateVerifyAccount(ctx context.Context, req repositories.CreateVerifyAccountRequest) error {
	verifyAccount := models.AccountVerify{
		Username:   req.Username,
		SecretCode: req.SecretCode,
	}
	return s.pg.Create(&verifyAccount).Error
}

func (s service) GetVerifyAccount(ctx context.Context, req repositories.GetVerifyAccountRequest) (repositories.GetVerifyAccountReply, error) {
	verifyAccount := models.AccountVerify{}
	if err := s.pg.Where("username = ?", req.Username).Take(&verifyAccount).Error; err != nil {
		return repositories.GetVerifyAccountReply{}, err
	}
	return repositories.GetVerifyAccountReply{AccountVerify: verifyAccount}, nil
}
