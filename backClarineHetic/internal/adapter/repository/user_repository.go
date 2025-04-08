package repository

import (
    "backClarineHetic/internal/domain"
    "database/sql"
    "errors"
    "github.com/google/uuid"
)

type PostgresUserRepo struct {
    db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
    return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(user *domain.User) error {
    user.UUID = uuid.New()
    query := `INSERT INTO users (uuid, username, email, password) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(query, user.UUID, user.Username, user.Email, user.Password)
    return err
}

func (r *PostgresUserRepo) FindByEmail(email string) (*domain.User, error) {
    query := `SELECT uuid, username, email, password FROM users WHERE email = $1`
    row := r.db.QueryRow(query, email)
    user := &domain.User{}
    if err := row.Scan(&user.UUID, &user.Username, &user.Email, &user.Password); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("utilisateur non trouvé")
        }
        return nil, err
    }
    return user, nil
}
