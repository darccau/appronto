package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Appointments AppointmentModel
	Users        UserModel
}

// Appointments interface {
// 	Insert(user *Appointment) error
// 	Get(id int64) (*Appointment, error)
// 	Update(user *Appointment) error
// 	Delete(id int64) error
// 	GetAll(email string, ilter Filters) ([]*Appointment, Metadata, error)
// }

// Users interface {
// 	Insert(user *User) error
// 	Get(id int64) (*User, error)
// 	Update(user *User) error
// 	Delete(id int64) error
// 	GetAll(email string, ilter Filters) ([]*User, Metadata, error)
// }

func NewModels(db *sql.DB) Models {
	return Models{
		Users:        UserModel{DB: db},
		Appointments: AppointmentModel{DB: db},
	}
}
