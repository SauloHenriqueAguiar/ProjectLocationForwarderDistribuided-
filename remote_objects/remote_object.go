package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Pega a porta a partir dos argumentos (se disponível)
	port := ":8081" // Porta padrão
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Inicia o listener na porta especificada
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Erro ao iniciar o objeto remoto:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Objeto remoto escutando na porta", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	locationRequest, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler requisição:", err)
		return
	}

	fmt.Printf("Objeto remoto recebeu a requisição: %s\n", locationRequest)

	response := fmt.Sprintf("Objeto remoto na porta %s processou a requisição: %s", conn.LocalAddr().String(), locationRequest)
	conn.Write([]byte(response))
}
