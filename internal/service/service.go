package service

import (
	"context"
	"log"

	"github.com/acool-kaz/post-crud-service-server/internal/models"
	"github.com/acool-kaz/post-crud-service-server/internal/repository"
)

type Post interface {
	Create(ctx context.Context, post models.Post) error
	Read(ctx context.Context) ([]models.Post, error)
	Update(ctx context.Context, id int, update models.UpdatePost) error
	Delete(ctx context.Context, id int) error
}

type Service struct {
	Post Post
}

func InitService(repo *repository.Repository) *Service {
	log.Println("init service")
	return &Service{
		Post: newPostService(repo.Post),
	}
}
