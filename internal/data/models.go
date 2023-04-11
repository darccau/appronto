package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")

	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Appointments AppointmentModel
	Users        UserModel
	Tokens       TokenModel
  Permissions PermissionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:        UserModel{DB: db},
		Appointments: AppointmentModel{DB: db},
		Tokens:       TokenModel{DB: db},
    Permissions:  PermissionsModel{DB: db},
	}
}
