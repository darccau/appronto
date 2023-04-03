package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/darccau/appronto/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64    `json:"id"`
	CreatedAt string   `json:"created_at"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	Activated bool     `json:"activated"`
	Version   int      `json:"-"`
}

type password struct {
	plaintext *string
	hash      []byte
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `
	 INSERT INTO users(first_name, last_name, password, email, activated)
	 VALUES ($1, $2, $3, $4, $5)
	 RETURNING id, created_at, version`

	args := []any{user.FirstName, user.LastName, user.Password.hash, user.Email, user.Activated}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (u UserModel) GetByEmail(email string) (*User, error) {
	var user User

	query := `
  SELECT id, created_at, first_name, last_name, password, email, activated, version
  FROM users
  WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
		&user.Password.hash,
		&user.Email,
		&user.Activated,
		&user.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Update(user *User) error {
	query := `
  UPDATE users
  SET first_name = $1, last_name = $2, password = $3, activated = $4, email = $5, version = version + 1
  WHERE id = $6 AND version = $7
  RETURNING version`

	args := []any{
		user.FirstName,
		user.LastName,
		user.Password.hash,
		user.Activated,
		user.Email,
		user.Id,
		user.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// TODO check if it's work
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case errors.Is(err, sql.ErrNoRows):
			// return ErrEditConflict
			return err
		default:
			return err
		}
	}
	return nil
}

func (u UserModel) GetAll(email string, filters Filters) ([]*User, Metadata, error) {
	query := fmt.Sprintf(`
    SELECT count(*) over(),id, first_name, last_name, email, password
    FROM users
    WHERE (to_tsvector('simple', email) @@ plainto_tsquery('simple', $1) OR $1 = '')
    ORDER BY %s %s, id ASC
    LIMIT $2 OFFSET $3
    `, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{email, filters.limit(), filters.offset()}

	rows, err := u.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	users := []*User{}

	for rows.Next() {
		var user User

		err := rows.Scan(
			&totalRecords,
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, Metadata{}, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return users, metadata, nil
}

func (u UserModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
  DELETE FROM users
  WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := u.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `
  SELECT users.id, users.created_at, users.first_name, users.last_name, users.email, users.password, users.activated, users.version
  FROM users
  INNER JOIN tokens
  ON users.id = tokens.user_id
  WHERE tokens.hash = $1
  AND tokens.scope = $2
  AND tokens.expiry > $3`

	args := []any{tokenHash[:], tokenScope, time.Now()}

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Id,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password.hash,
		&user.Activated,
		&user.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	// TODO check validate email regex
	// v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 characters long")
	v.Check(len(password) <= 72, "password", "must no be more than 72 characters long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.FirstName != "", "first name", "must be provided")
	v.Check(len(user.FirstName) <= 100, "first name", "must no be more than 100 characters")

	v.Check(user.LastName != "", "last name", "must be provided")
	v.Check(len(user.LastName) <= 100, "last name", "must no be more than 100 characters")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
		panic("missing password hash for users")
	}
}
