package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
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
  RETURNING id
  `

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
  WHERE id = $1
  `
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
