package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"
    "github.com/google/uuid"
)

type PostgresEventRepo struct {
    db *sql.DB
}

func NewPostgresEventRepo(db *sql.DB) *PostgresEventRepo {
    return &PostgresEventRepo{db: db}
}

func (r *PostgresEventRepo) Create(event *domain.Event) error {
    event.UUID = uuid.New()
    query := `INSERT INTO events (uuid, name, longitude, latitude, adress, city, start_date, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
    _, err := r.db.Exec(query, event.UUID, event.Name, event.Longitude, event.Latitude, event.Adress, event.City, event.StartDate, event.UserID)
    return err
}

func (r *PostgresEventRepo) GetByID(id uuid.UUID) (*domain.Event, error) {
    query := `SELECT uuid, name, longitude, latitude, adress, city, start_date, user_id FROM events WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    event := &domain.Event{}
    if err := row.Scan(&event.UUID, &event.Name, &event.Longitude, &event.Latitude, &event.Adress, &event.City, &event.StartDate, &event.UserID); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("événement non trouvé")
        }
        return nil, err
    }
    return event, nil
}

func (r *PostgresEventRepo) Update(event *domain.Event) error {
    query := `UPDATE events SET name = $1, longitude = $2, latitude = $3, adress = $4, city = $5, start_date = $6, user_id = $7 WHERE uuid = $8`
    _, err := r.db.Exec(query, event.Name, event.Longitude, event.Latitude, event.Adress, event.City, event.StartDate, event.UserID, event.UUID)
    return err
}

func (r *PostgresEventRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM events WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}
