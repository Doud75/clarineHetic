package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "github.com/google/uuid"
)

type PostgresMessageRepo struct {
    db *sql.DB
}

func NewPostgresMessageRepo(db *sql.DB) *PostgresMessageRepo {
    return &PostgresMessageRepo{db: db}
}

func (r *PostgresMessageRepo) Create(message *domain.Message) error {
    message.UUID = uuid.New()
    query := `INSERT INTO message (uuid, content, insert_at, user_id) VALUES ($1, $2, now(), $3)`
    _, err := r.db.Exec(query, message.UUID, message.Content, message.UserID)
    return err
}
