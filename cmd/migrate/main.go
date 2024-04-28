package main

import (
	"database/sql"
	"fmt"
	"go-core/constant"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const (
	migrationsDir = "databases/migrations"
)

func getDatabaseInstance() *sql.DB {
	env := constant.GetEnv()
	sqlConnUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	db, err := sql.Open("postgres", sqlConnUrl)
	if err != nil {
		panic(err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	return db
}

func Up() error {
	db := getDatabaseInstance()
	defer db.Close()
	return goose.Up(db, migrationsDir)
}

func main() {
	db := getDatabaseInstance()
	defer db.Close()
	var err error
	arguments := os.Args
	if len(arguments) <= 1 {
		log.Fatal("missing argument up or down")
	}
	log.Println("run ne", arguments)
	if arguments[1] == "up" {
		err = goose.Up(db, migrationsDir)

	} else if arguments[1] == "down" {
		err = goose.Down(db, migrationsDir)
	} else if arguments[1] == "create" {
		err = goose.Create(db, migrationsDir, arguments[2], "sql")
	} else if arguments[1] == "status" {
		err = goose.Status(db, migrationsDir)
	}
	if err != nil {
		panic(err)
	}
}
