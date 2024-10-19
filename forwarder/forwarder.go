package forwarder

import (
	"fmt"
	"net"
	"sync"
)

type RemoteObject struct {
	ID       string
	Location string
}

type Forwarder struct {
	remoteObjects []RemoteObject
	currentIndex  int
	mu            sync.Mutex
}

func NewForwarder(remoteObjects []RemoteObject) *Forwarder {
	return &Forwarder{
		remoteObjects: remoteObjects,
		currentIndex:  0,
	}
}

func (f *Forwarder) ForwardRequest(location string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Seleciona o próximo objeto remoto usando round-robin
	remoteObject := f.remoteObjects[f.currentIndex]
	f.currentIndex = (f.currentIndex + 1) % len(f.remoteObjects)

	// Encaminha a requisição para o objeto remoto selecionado
	conn, err := net.Dial("tcp", remoteObject.Location)
	if err != nil {
		return "", fmt.Errorf("não foi possível conectar ao objeto remoto: %v", err)
	}
	defer conn.Close()

	// Enviar a localização para o objeto remoto
	fmt.Fprintf(conn, location+"\n")

	// Ler a resposta do objeto remoto
	var response string
	fmt.Fscanf(conn, "%s\n", &response)

	return fmt.Sprintf("Objeto remoto na porta %s processou a requisição: %s", remoteObject.Location, response), nil
}
