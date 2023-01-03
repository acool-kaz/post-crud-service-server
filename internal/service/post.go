package service

import (
	"context"
	"fmt"

	"github.com/acool-kaz/post-crud-service-server/internal/models"
	"github.com/acool-kaz/post-crud-service-server/internal/repository"
)

type PostService struct {
	postRepo repository.Post
}

func newPostService(postRepo repository.Post) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) Create(ctx context.Context, post models.Post) error {
	if err := s.postRepo.Create(ctx, post); err != nil {
		return fmt.Errorf("post service: create: %w", err)
	}
	return nil
}

func (s *PostService) Read(ctx context.Context) ([]models.Post, error) {
	posts, err := s.postRepo.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("post service: read: %w", err)
	}
	return posts, nil
}

func (s *PostService) Update(ctx context.Context, id int, update models.UpdatePost) error {
	if err := s.postRepo.Update(ctx, id, update); err != nil {
		return fmt.Errorf("post service: update: %w", err)
	}
	return nil
}

func (s *PostService) Delete(ctx context.Context, id int) error {
	if err := s.postRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("post service: delete: %w", err)
	}
	return nil
}
