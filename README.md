# DougAPI
> 🇧🇷 [Versão em Português](#dougapi-em-português)  
> 🇺🇸 [English Version](#dougapi-in-english)

---

## **DougAPI em Português**
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

---

## **DougAPI in English**
DougAPI is a backend application developed in Go to manage an item control system using PostgreSQL. The project is modularly structured and follows best practices to facilitate maintenance and extensibility.

## **Table of Contents**
- [Technologies Used](#technologies-used)
- [Environment Setup](#environment-setup)
- [Running the Project](#running-the-project)
- [Project Structure](#project-structure)
- [Available Endpoints](#available-endpoints)
- [Contribution](#contribution)

## **Technologies Used**
- Language: [Go](https://go.dev/)
- Database: PostgreSQL
- Web Framework: [Echo](https://echo.labstack.com/)
- Connection Management: [pgx](https://github.com/jackc/pgx)
- Environment Variables Loader: [godotenv](https://github.com/joho/godotenv)

## **Environment Setup**
1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd DougAPI
   ```

2. **Configure environment variables:**
   Create a `.env` file in the project root with the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=database_name
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Set up the database:**
   Ensure that PostgreSQL is running and create the database specified in the `.env` file.

## **Running the Project**
- To run the application locally:
  ```bash
  go run cmd/api/main.go
  ```

- To run the tests (after implementation):
  ```bash
  make test
  ```

## **Project Structure**
```
DougAPI
|-- cmd
|   |-- api
|       |-- main.go          # Application entry point
|-- internal
|   |-- database
|   |   |-- postgres.go      # Database configuration
|   |-- handlers
|   |   |-- items.go         # Handlers for item endpoints
|   |-- models
|       |-- item.go          # Data structure for items
|-- .env                     # Environment variables (ignored by Git)
|-- go.mod                   # Dependency file
|-- go.sum                   # Dependency checksum
```

## **Available Endpoints**

### **GET /items**
Retrieves all items stored in the database.

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

## **Contribution**
Contributions are welcome! To contribute:
1. Fork the repository
2. Create a branch for your changes: `git checkout -b my-feature`
3. Commit your changes: `git commit -m 'My new feature'`
4. Push to your branch: `git push origin my-feature`
5. Open a Pull Request

---

This document will be updated as more features are implemented.



