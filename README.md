# Project Location Forwarder Distributed

Padrão de Projeto de Sistemas Distribuídos Location Forwarder

## Arquitetura

A arquitetura do projeto é composta por três componentes principais:

1. **Cliente**: Envia coordenadas no formato `latitude,longitude` para o servidor principal.
2. **Servidor Principal**: Recebe as coordenadas do cliente e encaminha a requisição para um dos objetos remotos, dependendo da lógica de balanceamento de carga.
3. **Objetos Remotos**: São servidores simples que recebem a requisição do servidor principal e retornam uma resposta confirmando o processamento das coordenadas.

### Fluxo de Requisição

1. O cliente envia uma localização (latitude, longitude) ao servidor principal.
2. O servidor principal escolhe um dos objetos remotos (porta 8081, 8082, 8083) e encaminha a requisição.
3. O objeto remoto processa a requisição e responde ao servidor principal.
4. O servidor principal retorna a resposta ao cliente.

## Pré-requisitos

- Go instalado na máquina. Para instalar, siga as instruções em: https://golang.org/doc/install
- Múltiplas instâncias do objeto remoto rodando em diferentes portas (`8081`, `8082`, `8083`).

## Como Executar o Projeto

### Passo 1: Clonar o Repositório

```bash
git clone https://github.com/SeuUsuario/ProjectLocationForwarderDistribuided.git
cd ProjectLocationForwarderDistribuided


### Passo 2: Executar os Objetos Remotos
 Execute três instâncias do objeto remoto em diferentes portas (8081, 8082, 8083).

# Instância 1 (porta 8081)
go run remote_object.go :8081

# Instância 2 (porta 8082)
go run remote_object.go :8082

# Instância 3 (porta 8083)
go run remote_object.go :8083


### Passo 3: Executar o Servidor Principal

Em uma nova janela de terminal, execute o servidor principal:

go run server.go

O servidor escutará na porta 8080 e encaminhará as requisições para os objetos remotos.

### Passo 4: Executar o Cliente
Agora, em uma nova janela de terminal, execute o cliente:

go run client.go

O cliente pedirá que você digite uma localização no formato latitude,longitude. Exemplo:

Digite a localização no formato latitude,longitude (ou 'sair' para encerrar): 12.34,56.78

O cliente enviará essa requisição ao servidor principal, que encaminhará a mensagem a um dos objetos remotos. A resposta será exibida no terminal do cliente.

Passo 5: Enviar Localizações
Digite diferentes localizações no terminal do cliente e observe o encaminhamento de requisições entre o servidor e os objetos remotos.

Exemplo de execução:
Digite a localização no formato latitude,longitude (ou 'sair' para encerrar): 11,22
Resposta do servidor: Objeto remoto na porta localhost:8081 processou a requisição: 11,22

Finalizar
Para encerrar o cliente, basta digitar sair.


Estrutura do Projeto:
ProjectLocationForwarderDistribuided/
│
├── client/
│   └── client.go          # Código do cliente
│
├── server/
│   └── server.go          # Código do servidor principal
│
├── remote_objects/
│   └── remote_object.go   # Código dos objetos remotos
│
├── forwarder.go           # Módulo de forwarding
└── README.md              # Documentação do projeto









