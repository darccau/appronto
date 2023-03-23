package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Appointments AppointmentModel
	Users        UserModel
	Tokens       TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:        UserModel{DB: db},
		Appointments: AppointmentModel{DB: db},
		Tokens:       TokenModel{DB: db},
	}
}
