package main

import (
	"banco-golang/model"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(localhost:3306)/banco?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println("Iniciando conexão com o banco...")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("Conexão iniciada com sucesso!")

	if err != nil {
		panic(errors.New("Erro ao conectar no banco de dados!"))
	}

	connection, err := db.DB()

	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
		}
		fmt.Println("Conexão com o banco interrompida!")
	}(connection)

	err = db.AutoMigrate(&model.Conta{}, &model.Usuario{})

	if err != nil {
		panic(errors.New("Erro ao conectar no banco de dados!"))
	}

}
