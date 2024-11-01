package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	cfg := mysql.Config{

		User:                 "root",
		Passwd:               EnvString("PASSWORD", "guest"),
		Net:                  EnvString("NET", "tcp"),
		Addr:                 EnvString("ADDR", "127.0.0.1:3308"),
		DBName:               EnvString("DBName", "db"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	storage := NewMySQLStorage(cfg)
	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}
	apiServer := NewAPIServer(":3000", db)
	apiServer.Run()

}
