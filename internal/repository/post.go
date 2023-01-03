package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/acool-kaz/post-crud-service-server/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func newPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(ctx context.Context, post models.Post) error {
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, title, body) VALUES($1, $2, $3, $4);", postTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("post repository: insert: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, post.Id, post.UserId, post.Title, post.Body); err != nil {
		return fmt.Errorf("post repository: insert: %w", err)
	}

	return nil
}

func (r *PostRepository) Read(ctx context.Context) ([]models.Post, error) {
	args := []interface{}{}

	query := fmt.Sprintf("SELECT * FROM %s;", postTable)

	id := ctx.Value(models.PostId)
	if id != nil {
		query = strings.ReplaceAll(query, ";", " WHERE id = $1;")
		args = append(args, id.(int))
	}

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("post repository: read: %w", err)
	}
	defer prep.Close()

	row, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("post repository: read: %w", err)
	}
	defer row.Close()

	var (
		posts []models.Post
		post  models.Post
	)

	for row.Next() {
		if err = row.Scan(&post.Id, &post.UserId, &post.Title, &post.Body); err != nil {
			return nil, fmt.Errorf("post repository: read: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) Update(ctx context.Context, id int, update models.UpdatePost) error {
	setParams := []string{}
	args := []interface{}{}
	argId := 1

	if update.Id > 0 {
		setParams = append(setParams, fmt.Sprintf("id=$%d", argId))
		args = append(args, update.Id)
		argId++
	}

	if update.UserId > 0 {
		setParams = append(setParams, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, update.UserId)
		argId++
	}

	if update.Title != "" {
		setParams = append(setParams, fmt.Sprintf("title=$%d", argId))
		args = append(args, update.Title)
		argId++
	}

	if update.Body != "" {
		setParams = append(setParams, fmt.Sprintf("body=$%d", argId))
		args = append(args, update.Body)
		argId++
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d;", postTable, strings.Join(setParams, ", "), argId)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("post repository: update: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, args...); err != nil {
		return fmt.Errorf("post repository: update: %w", err)
	}

	return nil
}

func (r *PostRepository) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", postTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("post repository: delete: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("post repository: delete: %w", err)
	}

	return nil
}
