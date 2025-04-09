package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"

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
    query := `INSERT INTO messages (uuid, content, insert_at, user_id, conversation_id) VALUES ($1, $2, $3, $4, $5)`
    _, err := r.db.Exec(query, message.UUID, message.Content, message.InsertAt, message.UserID, message.ConversationId)
    return err
}

func (r *PostgresMessageRepo) GetByID(id uuid.UUID) (*domain.Message, error) {
    query := `SELECT uuid, content, insert_at, user_id, conversation_id FROM messages WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    message := &domain.Message{}
    if err := row.Scan(&message.UUID, &message.Content, &message.InsertAt, &message.UserID, &message.ConversationId); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("message non trouvé")
        }
        return nil, err
    }
    return message, nil
}

func (r *PostgresMessageRepo) GetMessagesByConversationID(conversationID uuid.UUID) ([]*domain.Message, error) {
    query := `SELECT uuid, content, insert_at, user_id, conversation_id FROM messages where conversation_id = $1`
    rows, err := r.db.Query(query, conversationID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var messages []*domain.Message
    for rows.Next() {
        var message domain.Message
        if err = rows.Scan(&message.UUID, &message.Content, &message.InsertAt, &message.UserID, &message.ConversationId); err != nil {
            return nil, err
        }
        messages = append(messages, &message)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return messages, nil
}

func (r *PostgresMessageRepo) Update(message *domain.Message) error {
    query := `UPDATE messages SET content = $1, insert_at = $2, user_id = $3, conversation_id = $4 WHERE uuid = $5`
    _, err := r.db.Exec(query, message.Content, message.InsertAt, message.UserID, message.ConversationId, message.UUID)
    return err
}

func (r *PostgresMessageRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM messages WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}
