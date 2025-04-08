package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"

    "github.com/google/uuid"
)

type PostgresInstrumentRepo struct {
    db *sql.DB
}

func NewPostgresInstrumentRepo(db *sql.DB) *PostgresInstrumentRepo {
    return &PostgresInstrumentRepo{db: db}
}

func (r *PostgresInstrumentRepo) Create(instr *domain.Instrument) error {
    instr.UUID = uuid.New()
    query := `INSERT INTO instruments (uuid, name) VALUES ($1, $2)`
    _, err := r.db.Exec(query, instr.UUID, instr.Name)
    return err
}

func (r *PostgresInstrumentRepo) GetByID(id uuid.UUID) (*domain.Instrument, error) {
    query := `SELECT uuid, name FROM instruments WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    instr := &domain.Instrument{}
    if err := row.Scan(&instr.UUID, &instr.Name); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("instrument non trouvé")
        }
        return nil, err
    }
    return instr, nil
}

func (r *PostgresInstrumentRepo) Update(instr *domain.Instrument) error {
    query := `UPDATE instruments SET name = $1 WHERE uuid = $2`
    _, err := r.db.Exec(query, instr.Name, instr.UUID)
    return err
}

func (r *PostgresInstrumentRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM instruments WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}
