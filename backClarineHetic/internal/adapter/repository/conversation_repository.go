package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"

    "github.com/google/uuid"
)

type PostgresConversationRepo struct {
    db *sql.DB
}

func NewPostgresConversationRepo(db *sql.DB) *PostgresConversationRepo {
    return &PostgresConversationRepo{db: db}
}

func (r *PostgresConversationRepo) Create(convo *domain.Conversation) error {
    convo.UUID = uuid.New()
    query := `INSERT INTO conversations (uuid, user_id_a, user_id_b) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, convo.UUID, convo.UserIDA, convo.UserIDB)
    return err
}

func (r *PostgresConversationRepo) GetByID(id uuid.UUID) (*domain.Conversation, error) {
    query := `SELECT uuid, user_id_a, user_id_b FROM conversations WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    convo := &domain.Conversation{}
    if err := row.Scan(&convo.UUID, &convo.UserIDA, &convo.UserIDB); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("conversation non trouvée")
        }
        return nil, err
    }
    return convo, nil
}

func (r *PostgresConversationRepo) Update(convo *domain.Conversation) error {
    query := `UPDATE conversations SET user_id_a = $1, user_id_b = $2 WHERE uuid = $3`
    _, err := r.db.Exec(query, convo.UserIDA, convo.UserIDB, convo.UUID)
    return err
}

func (r *PostgresConversationRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM conversations WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}
