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
