package services

import (
	"fmt"
	"time"

	"github.com/morkid/paginate"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB Main database connection
var DB *gorm.DB

// PG Pagination library
var PG *paginate.Pagination

// InitDatabase initialize database connection
func InitDatabase() {
	if nil == DB {
		db := dbConnect()
		if nil != db {
			DB = db
			PG = paginate.New(&paginate.Config{
				FieldSelectorEnabled: true,
			})
		}
	}
}

func dbConnect() *gorm.DB {
	logLevel := logger.Info

	switch viper.GetString("ENVIRONMENT") {
	case "staging":
		logLevel = logger.Error
	case "production":
		logLevel = logger.Silent
	}

	config := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   viper.GetString("DB_TABLE_PREFIX"),
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &config)

	if nil != err {
		panic(err)
	}

	if nil != db {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetConnMaxLifetime(time.Second * 5)
	}

	return db
}
