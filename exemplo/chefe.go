package main

import (
	"os/exec"
	"fmt"
)

func execute() {
	out, err := exec.Command("./run.sh").Output()
	if err != nil {
		fmt.Printf("deu erro: %s", err)
	} else {
		fmt.Println("\nCommand Successfully Executed")
	}
	
	output := string(out[:])
	fmt.Println(output)
}


func main() {
	fmt.Printf("empregado, crie o relatorio\n")
	execute() // executando o comando de rodar empregado.go
	
}