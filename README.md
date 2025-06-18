# api-server-go
Simple CRUD API with Go, Fiber, Gorm, MySQL, Nats (Jetstream) and Redis.

## Dependências
### Obrigatórias
- [Golang](https://go.dev/doc/install)
- [Redis](https://redis.io/downloads/)
- [Make](https://www.gnu.org/software/make/)

### Opcionais
- [Postman](https://www.postman.com/downloads/)


## Instalação
### 1. Acessando o Projeto: 
Clonar o projeto _**api-server-go**_ pelo terminal digitando:
```bash
git clone https://github.com/yrnThiago/api-server-go.git
```

### 2. Instalação das dependências:
Após clonar o projeto, efetue a instalação das dependências digitando no terminal:
```bash
go mod
```

### 3. Configurar variáveis de ambiente:
Para configurar as variáveis de ambiente, crie um arquivo com o nome `.env` e copie o conteúdo do arquivo `.env.example` e cole dentro dele, ou peça para alguém do time enviar o arquivo `.env`.


### 4. Rodar o projeto:
Para rodar o projeto, no terminal:
```bash
make run-server
```


## Documentação
Para acessar a documentação da API utilize o Postman e importe todos os arquivos dentro de `docs/postman`.


## Dicas
- Para pular a autenticação, vá ao arquivo `.env` e altere a variável `SKIP_AUTH` para `true`.
