package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/acool-kaz/post-crud-service-server/internal/models"
)

const postTable = "posts"

type Post interface {
	Create(ctx context.Context, post models.Post) error
	Read(ctx context.Context) ([]models.Post, error)
	Update(ctx context.Context, id int, update models.UpdatePost) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	Post Post
}

func InitRepository(db *sql.DB) *Repository {
	log.Println("init repository")
	return &Repository{
		Post: newPostRepository(db),
	}
}
