package handlers

import (
	"fmt"
	"net/http"
	"os"

	"DougAPI/internal/database"
	"DougAPI/internal/models"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	pg *database.PostgreSQL
}

func NewItemHandler(pgCon *database.PostgreSQL) *ItemHandler {
	return &ItemHandler{
		pg: pgCon,
	}
}

// Handler para listar itens
func (ih ItemHandler) GetItemsHandle(c echo.Context) error {
	items, err := ih.pg.GetItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar itens"})
	}
	return c.JSON(http.StatusOK, items)
}

// Handler para criar itens
func (ih ItemHandler) CreateItemHandle(c echo.Context) error {
	var item models.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o JSON"})
	}

	item, err := ih.pg.CreateItem(item.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar item"})
	}

	return c.JSON(http.StatusCreated, item)
}

// Handler para buscar um item específico por ID
func (ih ItemHandler) GetItemByIDHandle(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID é obrigatório"})
	}

	item, err := ih.pg.GetItemByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Item não encontrado"})
	}

	return c.JSON(http.StatusOK, item)
}

// Handler para atualizar parcialmente um item por ID
func (ih ItemHandler) PatchItemHandle(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID é obrigatório"})
	}

	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o JSON"})
	}

	if err := ih.pg.PatchItem(id, updates); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar item"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item atualizado com sucesso"})
}

// Handler para atualizar completamente um item por ID
func (ih ItemHandler) PutItemHandle(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID é obrigatório"})
	}

	var item models.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao decodificar o JSON"})
	}

	if err := ih.pg.PutItem(id, item); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar item"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item atualizado com sucesso"})
}

// Handler para excluir um item por ID
func (ih ItemHandler) DeleteItemHandle(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID é obrigatório"})
	}

	if err := ih.pg.DeleteItem(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao deletar item"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (ih ItemHandler) DownloadAllItems(c echo.Context) error {
	items, err := ih.pg.GetItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar itens"})
	}

	// Cria o arquivo temporário
	filePath := "all_items.txt"
	file, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar arquivo"})
	}
	defer file.Close()

	// Escreve os dados no arquivo
	for _, item := range items {
		_, err := file.WriteString(fmt.Sprintf("ID: %d, Name: %s, CreatedAt: %s\n", item.ID, item.Name, item.CreatedAt))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao escrever no arquivo"})
		}
	}

	// Retorna o arquivo como download
	return c.Attachment(filePath, "all_items.txt")
}
