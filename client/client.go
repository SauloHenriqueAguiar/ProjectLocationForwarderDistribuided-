package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Digite a localização no formato latitude,longitude (ou 'sair' para encerrar): ")
		scanner.Scan()
		location := scanner.Text()

		if strings.ToLower(location) == "sair" {
			break
		}

		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("Erro ao conectar ao servidor:", err)
			continue
		}

		_, err = conn.Write([]byte(location + "\n"))
		if err != nil {
			fmt.Println("Erro ao enviar a localização:", err)
			conn.Close()
			continue
		}

		// Lendo a resposta do servidor
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Erro ao ler resposta do servidor:", err)
			conn.Close()
			continue
		}

		fmt.Println("Resposta do servidor:", string(response[:n]))

		conn.Close()
	}
}
