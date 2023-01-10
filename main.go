package main

import (
	"banco-golang/model"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	conta := model.Conta{"12345", "Conta Corrente", 100.50}
	pessoa := model.Pessoa{conta, "Guilherme da Silva Querentino", "51944463810", "11975890780", "qrt", "guigui10"}

	conta2 := model.Conta{"123456", "Conta Corrente", 120.50}
	pessoa2 := model.Pessoa{conta2, "Joaquim da Silva Sauro", "42833352998", "11985432691", "joaquim", "123"}

	var login string
	var senha string

	var mapaLogin = map[int64]model.Pessoa{
		1: pessoa,
		2: pessoa2,
	}

	fmt.Println("========================BANCO DO QRT================================")
	fmt.Println("Digite seu login:")
	fmt.Scanln(&login)
	fmt.Println("Digite sua senha:")
	fmt.Scanln(&senha)

	for _, j := range mapaLogin {
		if j.Login == login && j.Senha == senha {
			fmt.Println("Login com sucesso!")
			os.Exit(0)
		}
	}

	fmt.Println("Login errado, tente novamente!")
	time.Sleep(time.Second * 5)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	main()

}
