package connection

import (
	"gorm.io/gorm"

	"github.com/quocbang/grpc-gateway/server/repositories"
)

type DB struct {
	db gorm.DB
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
	scheme string
}

type Options func(*option)

func WithScheme(scheme string) Options {
	return func(o *option) {
		o.scheme = scheme
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

func NewRepositories(db DataBaseConfig) repositories.Repositories {
	return DB{}
}
