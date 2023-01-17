package config

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

func RetornaConfigBanco() mysql.Config {
	return mysql.Config{
		User:                 os.Getenv("root"),
		Passwd:               os.Getenv("root"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "banco",
		AllowNativePasswords: true,
	}
}
