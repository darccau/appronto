package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Users interface {
		Insert(user *User) error
		Get(id int64) (*User, error)
		Update(user *User) error
		Delete(id int64) error
		GetAll(first_name string, last_name string, email string, ilter Filters) ([]*User, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: UserModel{DB: db},
	}
}

// func NewMockUser() Models {
// 	return Models{
// 		Users: MockUsersModel{},
// 	}
// }
