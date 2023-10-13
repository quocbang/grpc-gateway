package account

import (
	"context"

	"github.com/quocbang/grpc-gateway/pkg/pb"
	"github.com/quocbang/grpc-gateway/server/repositories"
	mocks "github.com/quocbang/grpc-gateway/server/repositories/mock"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"github.com/quocbang/grpc-gateway/server/tests/setup/dependency"
	"github.com/quocbang/grpc-gateway/server/utils/hashing"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"github.com/stretchr/testify/mock"
)

func (s *Suite) TestLogin() {
	ctx := context.Background()
	assertion := s.Assertions
	testUsername := "test_username"
	testPassword := "test_password"
	testEmail := "test@gmail.com"
	hashPassword, err := hashing.HashPassword(testPassword)
	assertion.NoError(err)

	// good case
	{
		// Arrange
		req := &pb.LoginRequest{
			Username: testUsername,
			Password: testPassword,
		}
		repo := mocks.Repositories{}
		repo.EXPECT().Account().ReturnArguments = mock.Arguments{
			func() repositories.Account {
				account := mocks.Account{}
				getAccountRequest := repositories.GetAccountRequest{
					Username: testUsername,
				}
				account.EXPECT().GetAccount(ctx, getAccountRequest).ReturnArguments = mock.Arguments{
					repositories.GetAccountReply{
						Account: models.Account{
							Username:       testUsername,
							Email:          testEmail,
							IsUserVerified: false,
							Password:       hashPassword,
							Role:           roles.Roles_UNSPECIFIED_USER,
						},
					}, nil,
				}
				return &account
			}(),
		}
		mockRepo := s.NewMockRepositories(&repo.Mock)
		opts := []dependency.ServerTestOptions{
			dependency.WithMockRepositories(mockRepo),
			dependency.WithMockAccessTokenDuration(100),
			dependency.WithMockRefreshTokenDuration(1000),
			dependency.WithMockSecretKey("test_secret_key"),
		}
		mockServer := dependency.NewMockServer(opts...)

		// Act
		reply, err := mockServer.Account.Login(ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotEmpty(reply.AccessToken)
		assertion.NotEmpty(reply.RefreshToken)
		assertion.NotEmpty(reply.SessionId)
	}

	// insufficiency request

	// wrong password
}

func (s *Suite) TestSignUp() {

}

func (s *Suite) TestRenewAccess() {

}
