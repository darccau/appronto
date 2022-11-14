package data

import (
	"database/sql"
	"errors"

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
	v.Check(len(user.FirstName) <= 50, "first_name", "most not be more than 500bytes long")

	v.Check(user.LastName != "", "last_name", "must be provided")
	v.Check(len(user.LastName) <= 50, "last_name", "most not be more than 500bytes long")

	v.Check(user.Password != "", "password", "must be provided")
	v.Check(len(user.Password) >= 8, "password", "must be greater than 8 characters")
	v.Check(len(user.Password) <= 50, "password", "must be less than 50 characters")

	v.Check(user.Email != "", "email", "must be provided")
	v.Check(len(user.Email) >= 20, "password", "must be greater than 20")
	v.Check(len(user.Email) <= 50, "password", "must be less than 50")
}

func (u UserModel) Insert(user *User) error {
	query := `
  INSERT INTO users(first_name, last_name, password, email)
  VALUES ($1, $2, $3, $4)
  RETURNING id
  `
	args := []any{user.FirstName, user.LastName, user.Password, user.Email}

	return u.DB.QueryRow(query, args...).Scan(&user.Id)
}

func (u UserModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
  SELECT id, first_name, last_name, password, email 
  FROM users
  WHERE id = $1
  `

	var user User

	err := u.DB.QueryRow(query, id).Scan(
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

func (u UserModel) Update(user *User) {
	query := `
  UPDATE users
  SET first_name = $1, last_name = $2, password = $3, email = $4
  where id = $5
  `
	args := []any{
		user.Id,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Email,
	}

	u.DB.QueryRow(query, args...)
}

func (u UserModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
  DELETE FROM users 
  WHERE id = $1`

	result, err := u.DB.Exec(query, id)
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
