package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/darccau/appronto/internal/validator"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type UserModel struct {
	DB *sql.DB
}

func ValidateUsers(v *validator.Validator, user *User) {
	v.Check(user.FirstName != "", "first_name", "must be provided")
	v.Check(len(user.FirstName) <= 50, "first_name", "most not be more than 500 characters long")

	v.Check(user.LastName != "", "last_name", "must be provided")
	v.Check(len(user.LastName) <= 50, "last_name", "most not be more than 500 characters long")

	v.Check(user.Password != "", "password", "must be provided")
	v.Check(len(user.Password) <= 50, "password", "must be less than 50 characters long")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) <= 50, "email", "must not be less than 50 characters long")
}

func (u UserModel) Insert(user *User) error {

	query := `
  INSERT INTO users(first_name, last_name, password, email)
  VALUES ($1, $2, $3, $4)
  RETURNING id
  `
	args := []any{user.FirstName, user.LastName, user.Password, user.Email}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return u.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
}

func (u UserModel) Get(id int64) (*User, error) {
	var user User

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
  SELECT id, first_name, last_name, password, email 
  FROM users
  WHERE id = $1
  `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := u.DB.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
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

func (u UserModel) GetAll(first_name string, last_name string, email string, filter Filters) ([]*User, error) {
	query := `
    SELECT id, first_name, last_name, email, password 
    FROM users
    WHERE (LOWER(email) = LOWER($1) OR $1 = '')
    AND (LOWER(first_name) = LOWER($2) OR $2 = '')
    AND (LOWER(last_name) = LOWER($3) OR $3 = '')
    ORDER BY id
	 `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := u.DB.QueryContext(ctx, query, email, first_name, last_name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {

		var user User

		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserModel) Update(user *User) error {
	query := `
  UPDATE users
  SET first_name = $1, last_name = $2, password = $3, email = $4
  where id = $5
  RETURNING id
  `

	args := []any{
		user.FirstName,
		user.LastName,
		user.Password,
		user.Email,
		user.Id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// TODO check if it's work
	err := u.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return err
		default:
			return err
		}
	}

	return nil
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
