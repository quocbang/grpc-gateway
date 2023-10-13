package dependency

import (
	"fmt"
	"log"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"gorm.io/gorm"

	"github.com/quocbang/grpc-gateway/server/repositories"
	"github.com/quocbang/grpc-gateway/server/repositories/connection"
)

type Postgres struct {
	Name     string `long:"postgres-database-name-test" description:"postgres database name" env:"PG_DATABASE_NAME_TEST"`
	Address  string `long:"postgres-address-test" description:"postgres address" env:"PG_ADDRESS_TEST"`
	Port     int    `long:"postgres-port-test" description:"postgres port" env:"PG_PORT_TEST"`
	UserName string `long:"postgres-username-test" description:"postgres username" env:"PG_USERNAME_TEST"`
	Password string `long:"postgres-password-test" description:"postgres password" env:"PG_PASSWORD_TEST"`
	Schema   string `long:"postgres-schema-test" description:"postgres schema" env:"PG_SCHEMA_TEST"`
}

var databaseConfig struct {
	postgres Postgres
}

func parseFlags() error {
	configuration := []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "postgres configuration",
			LongDescription:  "postgres  configuration",
			Options:          &databaseConfig.postgres,
		},
	}

	parse := flags.NewParser(nil, flags.IgnoreUnknown)
	for _, opt := range configuration {
		if _, err := parse.AddGroup(opt.LongDescription, opt.LongDescription, opt.Options); err != nil {
			return err
		}
	}

	if _, err := parse.Parse(); err != nil {
		return fmt.Errorf("failed to parse database flags")
	}

	return nil
}

func InitRepositoriesTest() (repositories.Repositories, error) {
	opts := []connection.Options{
		connection.WithSchema(databaseConfig.postgres.Schema),
		connection.WithAutoMigrate(),
	}
	repo, err := connection.NewRepositories(connection.DataBaseConfig{
		Postgres: connection.PostgresConfig{
			Address:  databaseConfig.postgres.Address,
			Port:     databaseConfig.postgres.Port,
			Name:     databaseConfig.postgres.Name,
			Username: databaseConfig.postgres.UserName,
			Password: databaseConfig.postgres.Password,
		},
	}, opts...)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

type DatabaseConnection struct {
	PG *gorm.DB
}

func InitDatabaseConnection() (*DatabaseConnection, error) {
	pg, err := connection.NewPostgresConnection(connection.PostgresConfig{
		Address:  databaseConfig.postgres.Address,
		Port:     databaseConfig.postgres.Port,
		Name:     databaseConfig.postgres.Name,
		Username: databaseConfig.postgres.UserName,
		Password: databaseConfig.postgres.Password,
	}, databaseConfig.postgres.Schema)
	if err != nil {
		return nil, fmt.Errorf("failed to init new postgres conn, error: %v", err)
	}

	return &DatabaseConnection{PG: pg}, nil
}

func init() {
	err := parseFlags()
	if err != nil {
		log.Fatalf("failed to parse flag, error: %v", err)
	}
}
