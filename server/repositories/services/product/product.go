package product

import (
	"context"
	"fmt"
	"reflect"

	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) repositories.Product {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, req models.Product) error {
	if !reflect.DeepEqual(req, models.Product{}) {
		return fmt.Errorf("missing product")
	}

	return r.db.Create(&req).Error
}

func (r *repo) Creates(ctx context.Context, reqs []models.Product) error {
	if len(reqs) == 0 {
		return fmt.Errorf("missing products")
	}

	return r.db.Create(&reqs).Error
}
