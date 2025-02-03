package database

import (
	"context"
	"fmt"
	"log"

	"DougAPI/internal/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQL struct {
	db      *pgxpool.Pool
	options *PostgreSQLOptions
}

func NewPostgreSQL(options *PostgreSQLOptions) *PostgreSQL {
	if options == nil {
		log.Fatalf("Nenhuma configuração encontrada")
		return nil
	}

	// Configurar pool de conexões
	db, err := pgxpool.Connect(context.Background(), options.dsn)
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		return nil
	}

	log.Println("Conectado ao banco de dados com sucesso!")
	return &PostgreSQL{
		db:      db,
		options: options,
	}
}

// Fecha o pool de conexões
func (pg *PostgreSQL) Close() {
	pg.db.Close()
}

// Busca todos os itens do banco
func (pg *PostgreSQL) GetItems() ([]models.Item, error) {
	var items []models.Item
	query := "SELECT id, name, created_at FROM items"

	// Usa pgxscan para mapear o resultado diretamente para uma slice de structs
	err := pgxscan.Select(context.Background(), pg.db, &items, query)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// Cria um novo item no banco
func (pg *PostgreSQL) CreateItem(name string) (models.Item, error) {
	var item models.Item
	query := "INSERT INTO items (name) VALUES ($1) RETURNING id, name, created_at"

	// Usa QueryRow para obter o novo item criado
	err := pg.db.QueryRow(context.Background(), query, name).Scan(&item.ID, &item.Name, &item.CreatedAt)
	if err != nil {
		return models.Item{}, err
	}

	item.Name = name
	return item, nil
}

// Busca um item específico pelo ID
func (pg *PostgreSQL) GetItemByID(id string) (models.Item, error) {
	var item models.Item
	query := "SELECT id, name, created_at FROM items WHERE id = $1"

	// Usa pgxscan para mapear o resultado diretamente para a struct
	err := pgxscan.Get(context.Background(), pg.db, &item, query, id)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

// Atualiza parcialmente um item
func (pg *PostgreSQL) PatchItem(id string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return fmt.Errorf("nenhum campo fornecido para atualização")
	}

	query := "UPDATE items SET "
	args := []interface{}{}
	i := 1

	for key, value := range updates {
		query += fmt.Sprintf("%s = $%d, ", key, i)
		args = append(args, value)
		i++
	}
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d", i)
	args = append(args, id)

	_, err := pg.db.Exec(context.Background(), query, args...)
	return err
}

// Atualiza completamente um item
func (pg *PostgreSQL) PutItem(id string, item models.Item) error {
	query := "UPDATE items SET name = $1 WHERE id = $2"
	_, err := pg.db.Exec(context.Background(), query, item.Name, id)
	return err
}

// Exclui um item por ID
func (pg *PostgreSQL) DeleteItem(id string) error {
	query := "DELETE FROM items WHERE id = $1"
	_, err := pg.db.Exec(context.Background(), query, id)
	return err
}
