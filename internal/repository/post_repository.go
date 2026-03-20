package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) domain.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	query := `
		INSERT INTO posts (title, description, content, user_id, category_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		post.Title,
		post.Description,
		post.Content,
		post.UserID,
		post.CategoryID,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("db.create.post: %w", err)
	}
	return post, nil
}

func (r *postRepository) FindByID(ctx context.Context, id int64) (*domain.Post, error) {
	query := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`
	post := &domain.Post{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Description,
		&post.Content,
		&post.UserID,
		&post.CategoryID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NewNotFoundError("post", "id", id)
		}
		return nil, fmt.Errorf("db.post.find_by_id: %w", err)
	}
	return post, nil
}

func (r *postRepository) FindAll(ctx context.Context) ([]*domain.Post, error) {
	query := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		ORDER BY id ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("db.post.find_all: %w", err)
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.UserID,
			&post.CategoryID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.post.find_all.scan: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.post.find_all.rows: %w", err)
	}
	return posts, nil
}

func (r *postRepository) FindByCategoryID(ctx context.Context, categoryID int64) ([]*domain.Post, error) {
	query := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		WHERE category_id = $1
		ORDER BY id ASC
	`
	rows, err := r.db.QueryContext(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("db.post.find_by_category_id: %w", err)
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.UserID,
			&post.CategoryID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.post.find_by_category_id.scan: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.post.find_by_category_id.rows: %w", err)
	}
	return posts, nil
}

func (r *postRepository) Search(ctx context.Context, query string) ([]*domain.Post, error) {
	searchPattern := "%" + query + "%"
	q := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		WHERE title ILIKE $1 OR description ILIKE $1 OR content ILIKE $1
		ORDER BY id ASC
	`
	rows, err := r.db.QueryContext(ctx, q, searchPattern)
	if err != nil {
		return nil, fmt.Errorf("db.post.search: %w", err)
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.UserID,
			&post.CategoryID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.post.search.scan: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.post.search.rows: %w", err)
	}
	return posts, nil
}

func (r *postRepository) Update(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	query := `
		UPDATE posts
		SET title = $1, description = $2, content = $3, user_id = $4, category_id = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`
	err := r.db.QueryRowContext(ctx, query,
		post.Title,
		post.Description,
		post.Content,
		post.UserID,
		post.CategoryID,
		post.ID,
	).Scan(&post.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NewNotFoundError("post", "id", post.ID)
		}
		return nil, fmt.Errorf("db.post.update: %w", err)
	}
	return post, nil
}

func (r *postRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("db.post.delete: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("db.post.delete: %w", err)
	}
	if rowsAffected == 0 {
		return errorx.NewNotFoundError("post", "id", id)
	}
	return nil
}

func (r *postRepository) FindAllPaginated(ctx context.Context, offset, limit int) (*domain.PaginatedPosts, error) {
	// Get total count
	countQuery := `SELECT COUNT(*) FROM posts`
	var total int64
	err := r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("db.post.find_all_paginated.count: %w", err)
	}

	// Get paginated data
	query := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("db.post.find_all_paginated: %w", err)
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.UserID,
			&post.CategoryID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.post.find_all_paginated.scan: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.post.find_all_paginated.rows: %w", err)
	}

	return &domain.PaginatedPosts{
		Posts: posts,
		Total: total,
	}, nil
}

func (r *postRepository) FindByCategoryIDPaginated(ctx context.Context, categoryID int64, offset, limit int) (*domain.PaginatedPosts, error) {
	// Get total count
	countQuery := `SELECT COUNT(*) FROM posts WHERE category_id = $1`
	var total int64
	err := r.db.QueryRowContext(ctx, countQuery, categoryID).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("db.post.find_by_category_id_paginated.count: %w", err)
	}

	// Get paginated data
	query := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		WHERE category_id = $1
		ORDER BY id DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, categoryID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("db.post.find_by_category_id_paginated: %w", err)
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.UserID,
			&post.CategoryID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.post.find_by_category_id_paginated.scan: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.post.find_by_category_id_paginated.rows: %w", err)
	}

	return &domain.PaginatedPosts{
		Posts: posts,
		Total: total,
	}, nil
}

func (r *postRepository) SearchPaginated(ctx context.Context, query string, offset, limit int) (*domain.PaginatedPosts, error) {
	searchPattern := "%" + query + "%"

	// Get total count
	countQuery := `
		SELECT COUNT(*) FROM posts
		WHERE title ILIKE $1 OR description ILIKE $1 OR content ILIKE $1
	`
	var total int64
	err := r.db.QueryRowContext(ctx, countQuery, searchPattern).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("db.post.search_paginated.count: %w", err)
	}

	// Get paginated data
	q := `
		SELECT id, title, description, content, user_id, category_id, created_at, updated_at
		FROM posts
		WHERE title ILIKE $1 OR description ILIKE $1 OR content ILIKE $1
		ORDER BY id DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, q, searchPattern, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("db.post.search_paginated: %w", err)
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.UserID,
			&post.CategoryID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.post.search_paginated.scan: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.post.search_paginated.rows: %w", err)
	}

	return &domain.PaginatedPosts{
		Posts: posts,
		Total: total,
	}, nil
}
