package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"ProjectLocationForwarderDistribuided/forwarder"
)

func main() {
	// Define os objetos remotos
	remoteObjects := []forwarder.RemoteObject{
		{ID: "Object1", Location: "localhost:8081"},
		{ID: "Object2", Location: "localhost:8082"},
		{ID: "Object3", Location: "localhost:8083"},
	}

	// Cria o forwarder com os objetos remotos
	fwd := forwarder.NewForwarder(remoteObjects)

	// Inicia o servidor na porta 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Servidor escutando na porta 8080...")

	for {
		// Aceita novas conexões
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleRequest(conn, fwd)
	}
}

func handleRequest(conn net.Conn, fwd *forwarder.Forwarder) {
	defer conn.Close()

	// Lê a requisição do cliente
	reader := bufio.NewReader(conn)
	locationRequest, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler requisição:", err)
		return
	}

	locationRequest = strings.TrimSpace(locationRequest)
	fmt.Printf("Recebendo requisição de localização: %s\n", locationRequest)

	// Encaminha a requisição para o objeto remoto através do forwarder
	response, err := fwd.ForwardRequest(locationRequest)
	if err != nil {
		fmt.Println("Erro ao encaminhar a requisição:", err)
		conn.Write([]byte("Erro ao encaminhar a requisição\n"))
		return
	}

	// Envia a resposta de volta para o cliente
	conn.Write([]byte(response + "\n"))
	fmt.Printf("Resposta enviada ao cliente: %s\n", response)
}
