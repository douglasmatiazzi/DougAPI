# DougAPI

DougAPI é uma aplicação backend desenvolvida em Go para gerenciar um sistema de controle de itens utilizando PostgreSQL. O projeto está estruturado de forma modular e segue boas práticas para facilitar a manutenção e extensibilidade.

## **Sumário**
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Configuração do Ambiente](#configuração-do-ambiente)
- [Execução do Projeto](#execução-do-projeto)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Endpoints Disponíveis](#endpoints-disponíveis)
- [Contribuição](#contribuição)

## **Tecnologias Utilizadas**
- Linguagem: [Go](https://go.dev/)
- Banco de Dados: PostgreSQL
- Framework Web: [Echo](https://echo.labstack.com/)
- Gerenciamento de Conexão: [pgx](https://github.com/jackc/pgx)
- Carregamento de Variáveis de Ambiente: [godotenv](https://github.com/joho/godotenv)

## **Configuração do Ambiente**
1. **Clone o repositório:**
   ```bash
   git clone <url-do-repositorio>
   cd DougAPI
   ```

2. **Configure as variáveis de ambiente:**
   Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   DB_NAME=nome_do_banco
   ```

3. **Instale as dependências:**
   ```bash
   go mod tidy
   ```

4. **Configure o banco de dados:**
   Certifique-se de que o PostgreSQL esteja rodando e crie o banco de dados especificado no arquivo `.env`.

## **Execução do Projeto**
- Para rodar a aplicação localmente:
  ```bash
  go run cmd/api/main.go
  ```

- Para rodar os testes (após implementação):
  ```bash
  make test
  ```

## **Estrutura do Projeto**
```
DougAPI
|-- cmd
|   |-- api
|       |-- main.go          # Ponto de entrada da aplicação
|-- internal
|   |-- database
|   |   |-- postgres.go      # Configuração do banco de dados
|   |-- handlers
|   |   |-- items.go         # Handlers para os endpoints de itens
|   |-- models
|       |-- item.go          # Estrutura de dados dos itens
|-- .env                     # Variáveis de ambiente (ignorado pelo Git)
|-- go.mod                   # Arquivo de dependências
|-- go.sum                   # Checksum das dependências
```

## **Endpoints Disponíveis**

### **GET /items**
Lista todos os itens armazenados no banco de dados.

- **Request:**
  ```bash
  GET /items
  ```

- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "Item 1",
      "created_at": "2025-02-03T10:00:00Z"
    },
    {
      "id": 2,
      "name": "Item 2",
      "created_at": "2025-02-03T10:05:00Z"
    }
  ]
  ```

Esse documento será atualizado conforme mais funcionalidades forem implementadas.

