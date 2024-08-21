package connection

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/connection/logging"
	"github.com/quocbang/grpc-gateway/server/repositories/orm/models"
)

type DB struct {
	Postgres *gorm.DB
	TxFlag   bool
}

type PostgresConfig struct {
	Address  string
	Port     int
	Name     string
	Username string
	Password string
}

type DataBaseConfig struct {
	Postgres PostgresConfig
}

// option is save option with multiple option.
type option struct {
	scheme      string
	withMigrate bool
}

type Options func(*option)

func WithSchema(scheme string) Options {
	return func(o *option) {
		o.scheme = scheme
	}
}

func WithAutoMigrate() Options {
	return func(o *option) {
		o.withMigrate = true
	}
}

// parseOptions is parse option into connection.option struct.
func parseOptions(opts ...Options) *option {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func NewPostgresConnection(cfs PostgresConfig, scheme string) (*gorm.DB, error) {
	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		cfs.Address,
		cfs.Username,
		cfs.Password,
		cfs.Name,
		cfs.Port,
	)
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   scheme,
			SingularTable: false,
		},
		Logger: logging.NewGormLogger(),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewRepositories(dbCfs DataBaseConfig, options ...Options) (repositories.Repositories, error) {
	opts := parseOptions(options...)

	db, err := NewPostgresConnection(dbCfs.Postgres, opts.scheme)
	if err != nil {
		return nil, fmt.Errorf("failed to init new Postgres connection, error: %v", err)
	}

	if opts.withMigrate {
		if err := autoMigrate(db); err != nil {
			return nil, err
		}
	}

	return DB{
		Postgres: db,
	}, nil
}

func autoMigrate(db *gorm.DB) error {
	ms := models.ListModels()

	dst := make([]interface{}, len(ms))
	for i, m := range ms {
		dst[i] = m
	}

	return db.AutoMigrate(dst...)
}
