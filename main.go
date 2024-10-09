package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		// to-do : shift to env var
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "localhost.3306",
		DBName:               "pensieve",
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
