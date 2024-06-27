package db

import (
	"fmt"
	"time"

	"github.com/familybook-project/familybook-api-gin/src/pkg/util"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgresql struct {
	DB *gorm.DB
}

var DBInstance *Postgresql

func Connect() (*Postgresql, error) {
	err := util.ReadConfig()
	if err != nil {
		return nil, err
	}
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	dbname := viper.GetString("db.name")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	sslmode := viper.GetString("db.sslmode")
	timeZone := viper.GetString("db.timeZone")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(3 * time.Minute)

	return &Postgresql{DB: db}, nil
}
