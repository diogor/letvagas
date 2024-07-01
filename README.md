# Grupo Let Vagas

## Requisitos de desenvolvimento
- Golang 1.21+
- *make (opcional)*
- *Docker (opcional)*

## Configuração e instalação
- Copiar `.env.example` -> `.env` e editar de acordo com o seu ambiente
- Rode o script `database/init_db.sql` no seu banco de dados

## Build e execução
- Make: `make build`
- Manual (consulte o Makefile):
    ```bash
	go build -o ./build/letvagas
	cp -r ./templates ./build/
	cp -r ./static ./build/
    ```
- Executar: `./build/letvagas`
- *Servidor de desenvolvimento*: `go run server.go`

## Docker
- Build: `docker build -t letvagas .`
- Executar: `docker compose up -d`
