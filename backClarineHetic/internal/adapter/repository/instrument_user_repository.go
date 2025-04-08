package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"

    "github.com/google/uuid"
)

type PostgresInstrumentUserRepo struct {
    db *sql.DB
}

func NewPostgresInstrumentUserRepo(db *sql.DB) *PostgresInstrumentUserRepo {
    return &PostgresInstrumentUserRepo{db: db}
}

func (r *PostgresInstrumentUserRepo) Create(iu *domain.InstrumentUser) error {
    iu.UUID = uuid.New()
    query := `INSERT INTO instrument_users (uuid, user_id, instrument_id) VALUES ($1, $2, $3)`
    _, err := r.db.Exec(query, iu.UUID, iu.UserID, iu.InstrumentID)
    return err
}

func (r *PostgresInstrumentUserRepo) GetByID(id uuid.UUID) (*domain.InstrumentUser, error) {
    query := `SELECT uuid, user_id, instrument_id FROM instrument_users WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    iu := &domain.InstrumentUser{}
    if err := row.Scan(&iu.UUID, &iu.UserID, &iu.InstrumentID); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("instrument-user non trouvé")
        }
        return nil, err
    }
    return iu, nil
}

func (r *PostgresInstrumentUserRepo) Update(iu *domain.InstrumentUser) error {
    query := `UPDATE instrument_users SET user_id = $1, instrument_id = $2 WHERE uuid = $3`
    _, err := r.db.Exec(query, iu.UserID, iu.InstrumentID, iu.UUID)
    return err
}

func (r *PostgresInstrumentUserRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM instrument_users WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}
