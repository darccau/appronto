package data

import (
	"database/sql"
	"errors"
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

func (u UserModel) Update(user *User) error {
	return nil
}

func (u UserModel) Delete(id int64) error {
	return nil
}
