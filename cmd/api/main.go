package main

import (
	"log"
	"os"

	"DougAPI/internal/database"
	"DougAPI/internal/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Configurar o banco de dados
	pg := database.NewPostgreSQL(database.DefaultPostgreSQLOptions())
	if pg == nil {
		log.Fatalf("Erro ao conectar ao banco de dados")
	}
	defer pg.Close()

	// Inicializar Echo
	e := echo.New()

	// Configurar autenticação básica
	basicAuth := middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == os.Getenv("BASIC_AUTH_USER") && password == os.Getenv("BASIC_AUTH_PASS") {
			return true, nil
		}
		return false, nil
	})

	// Configurar Handlers
	handler := handlers.NewItemHandler(pg)

	// Rotas protegidas por Basic Auth
	g := e.Group("/items", basicAuth)
	g.GET("", handler.GetItemsHandle)
	g.GET("/:id", handler.GetItemByIDHandle)
	g.POST("", handler.CreateItemHandle)
	g.PATCH("/:id", handler.PatchItemHandle)
	g.PUT("/:id", handler.PutItemHandle)
	g.DELETE("/:id", handler.DeleteItemHandle)
	g.GET("/download", handler.DownloadAllItems)

	// Iniciar o servidor
	log.Println("Server is running on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
