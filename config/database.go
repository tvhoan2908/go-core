package config

import (
	"fmt"
	"go-core/constant"
	"log"

	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var once sync.Once

type Database struct {
	*gorm.DB
}

func GetDatabaseInstance() *gorm.DB {
	if DB == nil {
		once.Do(
			func() {
				log.Println("create single db instance now.")
				ConnectDatabase()
			})
	} else {
		log.Println("single db instance already created.")
	}

	return DB
}

func ConnectDatabase() *gorm.DB {
	env := constant.GetEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.DBHost, env.DBUsername, env.DBPassword, env.DBName, env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connect success to database: ", env.DBName)
	DB = db

	return db
}
