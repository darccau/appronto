package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/darccau/appronto/internal/data"
)

func (app *application) createAppointments(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DateTime time.Time `json:"date_time"`
		Reason   string    `json:"reason"`
		Notes    string    `json:"notes"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	appointment := &data.Appointment{
		DateTime: input.DateTime,
		Reason:   input.Reason,
		Notes:    input.Notes,
	}

	err = app.models.Appointments.Insert(appointment)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/appointments/%d", appointment.Id))

	err = app.writeJSON(w, http.StatusCreated, envelope{"appointment": appointment}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showAppointment(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	appointment, err := app.models.Appointments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"appointment": appointment}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAppointment(w http.ResponseWriter, r *http.Request) {
  id, err := app.readIdParam(r)
  if err != nil {
    app.notFoundResponse(w, r)
  }

  err = app.models.Appointments.Delete(id)
  if err != nil {
     switch {
     case errors.Is(err, data.ErrRecordNotFound):
        app.notFoundResponse(w,r)
     default:
        app.serverErrorResponse(w,r,err)
       }
    return
   }

  err = app.writeJSON(w, http.StatusOK, envelope{"message": "appointment was sucessfully deleted"}, nil)
  if err != nil {
    app.serverErrorResponse(w, r, err)
  }
}
