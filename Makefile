# Variáveis
APP_NAME=DougAPI

# Comando para rodar a aplicação
run:
	@echo "Iniciando $(APP_NAME)..."
	@go run cmd/api/main.go

# Comando para instalar/atualizar dependências
tidy:
	@echo "Organizando dependências..."
	@go mod tidy

# Comando para compilar o projeto
build:
	@echo "Compilando $(APP_NAME)..."
	@go build -o $(APP_NAME) cmd/api/main.go

# Limpar build
clean:
	@echo "Removendo binários..."
	@rm -f $(APP_NAME)

.PHONY: run tidy build clean
