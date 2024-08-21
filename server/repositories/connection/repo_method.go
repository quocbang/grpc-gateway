package connection

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/errors"
	"github.com/quocbang/grpc-gateway/server/repositories/services/account"
	"github.com/quocbang/grpc-gateway/server/repositories/services/product"
)

func (db DB) Account() repositories.Account {
	return account.NewAccount(db.Postgres)
}

func (db DB) Product() repositories.Product {
	return product.NewProduct(db.Postgres)
}

func (db DB) BeginTx(ctx context.Context, sqlOpts ...*sql.TxOptions) (repositories.Repositories, error) {
	if db.TxFlag {
		return nil, errors.Error{Code: errors.Code_ALREADY_IN_TRANSACTION}
	}
	newHandlerWithTransaction := DB{
		Postgres: db.Postgres.WithContext(ctx).Begin(sqlOpts...),
		TxFlag:   true,
	}
	return newHandlerWithTransaction, nil
}

func (db DB) Rollback() error {
	if !db.TxFlag {
		return fmt.Errorf("not in transaction")
	}
	return db.Postgres.Rollback().Error
}

func (db DB) Commit() error {
	if !db.TxFlag {
		return fmt.Errorf("not in transaction")
	}
	return db.Postgres.Commit().Error
}

func (db DB) Close() error {
	sqlDB, err := db.Postgres.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
