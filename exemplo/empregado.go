package main

import (
	"os/exec"
	"fmt"
)

func main() {
	terminal := exec.Command("touch", "relatorio.txt") // criando terminal e passando o comando
	fmt.Println("Ok, vou criar o relatorio")
	error := terminal.Run() // executando o comando de criar relatorio.txt
	
	if error != nil {
		fmt.Println("Erro ao executar comando: %s",error)
	} else {
		fmt.Println("Relatorio criado!")
	}
}