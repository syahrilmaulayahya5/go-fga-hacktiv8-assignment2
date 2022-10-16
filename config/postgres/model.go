package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
}

type PostgreClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func NewPostgresConnecion(config Config) PostgreClient {
	connectionString := fmt.Sprintf(`
		host=%s
		port=%s
		user=%s
		password=%s
		dbname=%s`,
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("error when connecting to database")
	}
	return &PostgresClientImpl{cln: db, config: config}
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}
