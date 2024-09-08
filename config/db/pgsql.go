package db

import (
	"fmt"
	"log"

	"github.com/go-prome/app/drivers/models"
	configs "github.com/go-prome/config"
	"github.com/go-prome/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client interface {
	Conn() *gorm.DB
	Close()
	WithTx(db *gorm.DB) Client
	Begin() *gorm.DB
}

func New(configuration configs.Config) Client {
	dsn := newPgSQLConfig(configuration).String()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.PanicIfError(err)
	//dinyalakan ketika diperlukan, jika tidak comment.
	err = db.AutoMigrate(&models.Note{})
	//exceptions.PanicIfError(err)

	// it'll handle based by : https://github.com/go-sql-driver/mysql/issues/674
	sqlDB, err := db.DB()
	utils.PanicIfError(err)
	sqlDB.SetConnMaxLifetime(0)
	log.Println("PgSql Framework Connected")
	return &client{db}
}

type client struct {
	db *gorm.DB
}

func (c client) Begin() *gorm.DB {
	return c.db.Begin()
}

func (c client) WithTx(db *gorm.DB) Client {
	c.db = db
	return &c
}

func (c *client) Conn() *gorm.DB {
	return c.db
}

func (c *client) Close() {
	sqlDB, err := c.db.DB()
	utils.PanicIfError(err)

	err = sqlDB.Close()
	utils.PanicIfError(err)
}

type pgSQLConfig struct {
	Host     string
	Password string
	Port     string
	User     string
	DBName   string
}

func newPgSQLConfig(configuration configs.Config) *pgSQLConfig {
	dbConfig := pgSQLConfig{
		Host:     configuration.Get("DB_HOST"),
		Port:     configuration.Get("DB_PORT"),
		User:     configuration.Get("DB_USER"),
		Password: configuration.Get("DB_PASSWORD"),
		DBName:   configuration.Get("DB_NAME"),
	}

	return &dbConfig
}

func (dbConfig *pgSQLConfig) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
