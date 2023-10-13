package email

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"github.com/quocbang/grpc-gateway/server/utils/hashing"
	"github.com/quocbang/grpc-gateway/server/utils/roles"
	"gorm.io/gorm"
)

func (s Suite) TestCreateAccount() {
	assertion := s.Assertions
	testUser := "test_user"
	testEmail := "test_email@gmail.com"
	testHashPassword, err := hashing.HashPassword("test_password")
	assertion.NoError(err)
	ctx := context.Background()

	// good case.
	{
		// Arrange
		assertion.NoError(s.ClearTable())
		req := repositories.CreateAccountRequest{
			Username:     testUser,
			Email:        testEmail,
			HashPassword: testHashPassword,
		}

		// Act
		err := s.Repo.Account().CreateAccount(ctx, req)

		// Assert
		assertion.NoError(err)
		reply, err := s.Repo.Account().GetAccount(ctx, repositories.GetAccountRequest{
			Username: testUser,
		})
		assertion.NoError(err)
		expected := repositories.GetAccountReply{
			Account: models.Account{
				Username:       testUser,
				Email:          testEmail,
				IsUserVerified: false,
				Password:       testHashPassword,
				Role:           roles.Roles_UNSPECIFIED_USER,
			},
		}
		assertion.Equal(expected, reply)
	}

	// bad case
	{ // account duplicated
		// Arrange
		req := repositories.CreateAccountRequest{
			Username:     testUser,
			Email:        testEmail,
			HashPassword: testHashPassword,
		}

		// Act
		err := s.Repo.Account().CreateAccount(ctx, req)

		// Assert
		assertion.Error(err)
		expected := fmt.Errorf("ERROR: duplicate key value violates unique constraint \"account_pkey\" (SQLSTATE 23505)")
		assertion.Equal(expected.Error(), err.Error())
	}
}

func (s *Suite) TestGetAccount() {
	assertion := s.Assertions
	testUsername := "test_user"
	testEmail := "test_email@gmail.com"
	testPassword := "test_password"
	ctx := context.Background()
	hashPassword, err := hashing.HashPassword(testPassword)
	assertion.NoError(err)

	// good case
	{
		// Arrange
		assertion.NoError(s.ClearTable())
		createAccountRequest := repositories.CreateAccountRequest{
			Username:     testUsername,
			Email:        testEmail,
			HashPassword: hashPassword,
		}
		err := s.Repo.Account().CreateAccount(ctx, createAccountRequest)
		assertion.NoError(err)

		req := repositories.GetAccountRequest{
			Username: testUsername,
		}

		// Act
		reply, err := s.Repo.Account().GetAccount(ctx, req)

		// Assert
		assertion.NoError(err)
		expected := repositories.GetAccountReply{
			Account: models.Account{
				Username:       testUsername,
				Email:          testEmail,
				IsUserVerified: false,
				Password:       hashPassword,
				Role:           roles.Roles_UNSPECIFIED_USER,
			},
		}
		assertion.Equal(expected, reply)
	}

	// base case
	{ // account does not exited
		// Arrange
		assertion.NoError(s.ClearTable())
		req := repositories.GetAccountRequest{
			Username: testUsername,
		}

		// Act
		_, err := s.Repo.Account().GetAccount(ctx, req)

		// Assert
		assertion.Error(err)
		assertion.Equal(gorm.ErrRecordNotFound, err)
	}
}

func (s *Suite) TestCreateVerifyAccount() {
	assertion := s.Assertions
	ctx := context.Background()
	testUsername := "test_username"
	testEmail := "test_email@gmail.com"
	testHashPassword, err := hashing.HashPassword("test_password")
	assertion.NoError(err)
	secretCode := uuid.NewString()

	// good case
	{
		// Arrange
		assertion.NoError(s.ClearTable())
		createAccountRequest := repositories.CreateAccountRequest{
			Username:     testUsername,
			Email:        testEmail,
			HashPassword: testHashPassword,
		}
		err := s.Repo.Account().CreateAccount(ctx, createAccountRequest)
		assertion.NoError(err)

		req := repositories.CreateVerifyAccountRequest{
			Username:   testUsername,
			SecretCode: secretCode,
		}

		// Act
		err = s.Repo.Account().CreateVerifyAccount(ctx, req)

		// Assert
		assertion.NoError(err)
		reply, err := s.Repo.Account().GetVerifyAccount(ctx, repositories.GetVerifyAccountRequest{
			Username: testUsername,
		})
		assertion.NoError(err)
		expected := repositories.GetVerifyAccountReply{
			AccountVerify: models.AccountVerify{
				Username:   testUsername,
				SecretCode: secretCode,
			},
		}
		assertion.Equal(expected.AccountVerify.Username, reply.AccountVerify.Username)
		assertion.Equal(expected.AccountVerify.SecretCode, reply.AccountVerify.SecretCode)
		assertion.NotNil(reply.AccountVerify.CreatedAt)
		assertion.NotNil(reply.AccountVerify.UpdatedAt)
	}

	// bad case
	{
		// Arrange
		assertion.NoError(s.ClearTable())
		req := repositories.CreateVerifyAccountRequest{
			Username:   testUsername,
			SecretCode: secretCode,
		}

		// Act
		err = s.Repo.Account().CreateVerifyAccount(ctx, req)

		// Assert
		assertion.Error(err)
		assertion.Equal("ERROR: insert or update on table \"account_verify\" violates foreign key constraint \"fk_account_account_verify\" (SQLSTATE 23503)", err.Error())
	}
}

func (s *Suite) TestGetVerifyAccount() {
	assertion := s.Assertions
	ctx := context.Background()
	testUsername := "test_username"
	testEmail := "test_email@gmail.com"
	testHashPassword, err := hashing.HashPassword("test_password")
	assertion.NoError(err)
	secretCode := uuid.NewString()

	// good case
	{
		// Arrange
		assertion.NoError(s.ClearTable())
		createAccountRequest := repositories.CreateAccountRequest{
			Username:     testUsername,
			Email:        testEmail,
			HashPassword: testHashPassword,
		}
		err := s.Repo.Account().CreateAccount(ctx, createAccountRequest)
		assertion.NoError(err)

		createVerifyAccountRequest := repositories.CreateVerifyAccountRequest{
			Username:   testUsername,
			SecretCode: secretCode,
		}
		err = s.Repo.Account().CreateVerifyAccount(ctx, createVerifyAccountRequest)
		assertion.NoError(err)
		req := repositories.GetVerifyAccountRequest{
			Username: testUsername,
		}

		// Act
		reply, err := s.Repo.Account().GetVerifyAccount(ctx, req)

		// Assert
		assertion.NoError(err)
		expected := repositories.GetVerifyAccountReply{
			AccountVerify: models.AccountVerify{
				Username:   testUsername,
				SecretCode: secretCode,
			},
		}
		assertion.Equal(expected.AccountVerify.Username, reply.AccountVerify.Username)
		assertion.Equal(expected.AccountVerify.SecretCode, reply.AccountVerify.SecretCode)
		assertion.NotNil(reply.AccountVerify.CreatedAt)
		assertion.NotNil(reply.AccountVerify.UpdatedAt)
	}
}
