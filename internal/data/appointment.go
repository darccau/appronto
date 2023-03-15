package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
	// "github.com/darccau/appronto/internal/validator"
)

type Appointment struct {
	Id         int64     `json:"id"`
	DateTime   time.Time `json:"date_time"`
	Reason     string    `json:"reason"`
	Notes      string    `json:"notes"`
	Created_at time.Time `json:"created_at"`
	Version    int64     `json:"version"`
}

type AppointmentModel struct {
	DB *sql.DB
}

func (a AppointmentModel) Insert(appointment *Appointment) error {
	query := `
  INSERT INTO appointments("date_time", "reason", "notes") 
  VALUES ($1, $2, $3)
  RETURNING id`

	args := []any{appointment.DateTime, appointment.Reason, appointment.Notes}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&appointment.Id)
}

func (a AppointmentModel) Get(id int64) (*Appointment, error) {
	var appointment Appointment

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
  SELECT id, date_time, reason, notes, create_at, version
  FROM appointments
  WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, id).Scan(
		&appointment.Id,
		&appointment.DateTime,
		&appointment.Reason,
		&appointment.Notes,
		&appointment.Created_at,
		&appointment.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &appointment, nil
}

func (a AppointmentModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
  DELETE FROM appointments
  WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := a.DB.ExecContext(ctx, query, id)
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

func (a AppointmentModel) Update(appointment *Appointment) error {
	query := `
	 UPDATE appointments
	 SET date_time = $1, reason = $2, notes = $3
	 WHERE id = $4
	 RETURNING id`

	args := []any{
		appointment.DateTime,
		appointment.Reason,
		appointment.Notes,
		appointment.Id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, args...).Scan(&appointment.Id)
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

// func ValidateUsers(v *validator.Validator, appointment *Appointment) {
// v.Check(user.FirstName != "", "first_name", "must be provided")
// v.Check(len(user.FirstName) <= 50, "first_name", "most not be more than 500 characters long")
//
// v.Check(user.LastName != "", "last_name", "must be provided")
// v.Check(len(user.LastName) <= 50, "last_name", "most not be more than 500 characters long")
//
// v.Check(user.Password != "", "password", "must be provided")
// v.Check(len(user.Password) <= 50, "password", "must be less than 50 characters long")
//
// v.Check(user.Email != "", "email", "must be provided")
// v.Check(len(user.Email) <= 50, "email", "must not be less than 50 characters long")
// }
