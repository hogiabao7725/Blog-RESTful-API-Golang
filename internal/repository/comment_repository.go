package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/hogiabao7725/blog-rest-api-golang/internal/domain"
	"github.com/hogiabao7725/blog-rest-api-golang/internal/errorx"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) domain.CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	query := `
		INSERT INTO comments (body, user_id, post_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := r.db.QueryRowContext(ctx, query, comment.Body, comment.UserID, comment.PostID).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("db.comment.create: %w", err)
	}

	return comment, nil
}

func (r *commentRepository) FindByID(ctx context.Context, id int64) (*domain.Comment, error) {
	query := `
		SELECT id, body, user_id, post_id, created_at
		FROM comments
		WHERE id = $1
	`

	comment := &domain.Comment{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&comment.ID,
		&comment.Body,
		&comment.UserID,
		&comment.PostID,
		&comment.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NewNotFoundError("comment", "id", id)
		}
		return nil, fmt.Errorf("db.comment.find_by_id: %w", err)
	}

	return comment, nil
}

func (r *commentRepository) FindByPostID(ctx context.Context, postID int64) ([]*domain.Comment, error) {
	query := `
		SELECT id, body, user_id, post_id, created_at
		FROM comments
		WHERE post_id = $1
		ORDER BY id ASC
	`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, fmt.Errorf("db.comment.find_by_post_id: %w", err)
	}
	defer rows.Close()

	var comments []*domain.Comment
	for rows.Next() {
		comment := &domain.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.Body,
			&comment.UserID,
			&comment.PostID,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("db.comment.find_by_post_id.scan: %w", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db.comment.find_by_post_id.rows: %w", err)
	}

	return comments, nil
}

func (r *commentRepository) Update(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	query := `
		UPDATE comments
		SET body = $1
		WHERE id = $2
		RETURNING post_id, user_id, created_at
	`

	err := r.db.QueryRowContext(ctx, query, comment.Body, comment.ID).Scan(
		&comment.PostID,
		&comment.UserID,
		&comment.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.NewNotFoundError("comment", "id", comment.ID)
		}
		return nil, fmt.Errorf("db.comment.update: %w", err)
	}

	return comment, nil
}

func (r *commentRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM comments
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("db.comment.delete: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("db.comment.delete.rows_affected: %w", err)
	}
	if rowsAffected == 0 {
		return errorx.NewNotFoundError("comment", "id", id)
	}

	return nil
}
