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

func (r *PostgresUserRepo) GetByID(id uuid.UUID) (*domain.User, error) {
    query := `SELECT uuid, username, email, password FROM users WHERE uuid = $1`
    row := r.db.QueryRow(query, id)
    user := &domain.User{}
    if err := row.Scan(&user.UUID, &user.Username, &user.Email, &user.Password); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("utilisateur non trouvé")
        }
        return nil, err
    }
    return user, nil
}

func (r *PostgresUserRepo) Update(user *domain.User) error {
    query := `UPDATE users SET username = $1, email = $2, password = $3 WHERE uuid = $4`
    _, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.UUID)
    return err
}

func (r *PostgresUserRepo) Delete(id uuid.UUID) error {
    query := `DELETE FROM users WHERE uuid = $1`
    _, err := r.db.Exec(query, id)
    return err
}

func (r *PostgresUserRepo) SearchProfiles(searchTerm string) ([]*domain.User, error) {
    query := `
		SELECT DISTINCT u.uuid, u.username, u.email, u.password
		FROM users u
		LEFT JOIN instrument_users iu ON u.uuid = iu.user_id
		LEFT JOIN instruments i ON i.uuid = iu.instrument_id
		WHERE u.username ILIKE '%' || $1 || '%'
		   OR i.name ILIKE '%' || $1 || '%'
	`
    rows, err := r.db.Query(query, searchTerm)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*domain.User
    for rows.Next() {
        var user domain.User
        if err = rows.Scan(&user.UUID, &user.Username, &user.Email, &user.Password); err != nil {
            return nil, err
        }
        users = append(users, &user)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return users, nil
}
