package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"

    "github.com/google/uuid"
)

type PostgresEventUserRepo struct {
    db *sql.DB
}

func NewPostgresEventUserRepo(db *sql.DB) *PostgresEventUserRepo {
    return &PostgresEventUserRepo{db: db}
}

func (r *PostgresEventUserRepo) Create(eu *domain.EventUser) error {
    eu.UUID = uuid.New()
    query := `INSERT INTO event_users (uuid, user_id, event_id) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, eu.UUID, eu.UserID, eu.EventID)
    return err
}

func (r *PostgresEventUserRepo) GetByID(id uuid.UUID) (*domain.EventUser, error) {
    query := `SELECT uuid, user_id, event_id FROM event_users WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    eu := &domain.EventUser{}
    if err := row.Scan(&eu.UUID, &eu.UserID, &eu.EventID); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("relation event-user non trouvée")
        }
        return nil, err
    }
    return eu, nil
}

func (r *PostgresEventUserRepo) Update(eu *domain.EventUser) error {
    query := `UPDATE event_users SET user_id = $1, event_id = $2 WHERE uuid = $3`
    _, err := r.db.Exec(query, eu.UserID, eu.EventID, eu.UUID)
    return err
}

func (r *PostgresEventUserRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM event_users WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}
