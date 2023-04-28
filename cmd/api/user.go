package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/darccau/appronto/internal/data"
	"github.com/darccau/appronto/internal/validator"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Activated: false,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.models.Permissions.AddForUser(user.Id, "appointments:read")
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	token, err := app.models.Tokens.New(user.Id, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		data := map[string]any{
			"activationToken": token.Plaintext,
			"UserId":          user.Id,
		}

		err = app.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			app.logger.PrintError(err, nil)
		}
	})

	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) activateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlaintext string `json:"token"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
	}

	user, err := app.models.Users.GetForToken(data.ScopeActivation, input.TokenPlaintext)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	user.Activated = true

	err = app.models.Users.Update(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.models.Tokens.DeleteAllForUsers(data.ScopeActivation, user.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
//
// 	var input struct {
// 		FirstName *string `json:"first_name"`
// 		LastName  *string `json:"last_name"`
// 		Password  *string `json:"password"`
// 		Email     *string `json:"email"`
// 	}
//
// 	fmt.Println(input)
//
// 	user, err := app.models.Users.GetByEmail()
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			app.notFoundResponse(w, r)
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}
//
// 	err = app.readJSON(w, r, &input)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}
//
// 	if input.FirstName != nil {
// 		user.FirstName = *input.FirstName
// 	}
//
// 	if input.LastName != nil {
// 		user.LastName = *input.LastName
// 	}
//
// 	if input.Password != nil {
// 		user.Password = *input.Password
// 	}
//
// 	if input.Email != nil {
// 		user.Email = *input.Email
// 	}
//
// 	v := validator.New()
//
// 	if data.ValidateUser(v, user); !v.Valid() {
// 		app.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}
//
// 	err = app.models.Users.Update(user)
//
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}
//
// 	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
//
// func (app *application) listUsers(w http.ResponseWriter, r *http.Request) {
//
// 	var input struct {
// 		FirstName string
// 		LastName  string
// 		Email     string
// 		data.Filters
// 	}
//
// 	v := validator.New()
// 	qs := r.URL.Query()
//
// 	input.FirstName = app.readString(qs, "first_name", "")
// 	input.LastName = app.readString(qs, "last_name", "")
// 	input.Email = app.readString(qs, "email", "")
//
// 	input.Filters.Page = app.readInt(qs, "page", 1, v)
// 	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
//
// 	input.Filters.Sort = app.readString(qs, "sort", "id")
// 	input.Filters.SortSafelist = []string{"id", "-id", "email", "-email", "page", "page_size"}
//
// 	if data.ValidateFilters(v, input.Filters); !v.Valid() {
// 		app.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}
//
// 	users, metadata, err := app.models.Users.GetAll(input.Email, input.Filters)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}
//
// 	err = app.writeJSON(w, http.StatusOK, envelope{"users": users, "metadata": metadata}, nil)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
//
// func (app *application) showUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := app.readIdParam(r)
// 	if err != nil {
// 		app.notFoundResponse(w, r)
// 		return
// 	}
// 	user, err := app.models.Users.Get(id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			app.notFoundResponse(w, r)
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}
//
// 	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
//
// func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := app.readIdParam(r)
// 	if err != nil {
// 		app.notFoundResponse(w, r)
// 		return
// 	}
//
// 	user, err := app.models.Users.Get(id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			app.notFoundResponse(w, r)
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}
//
// 	var input struct {
// 		FirstName *string `json:"first_name"`
// 		LastName  *string `json:"last_name"`
// 		Password  *string `json:"password"`
// 		Email     *string `json:"email"`
// 	}
//
// 	err = app.readJSON(w, r, &input)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}
//
// 	if input.FirstName != nil {
// 		user.FirstName = *input.FirstName
// 	}
//
// 	if input.LastName != nil {
// 		user.LastName = *input.LastName
// 	}
//
// 	if input.Password != nil {
// 		user.Password = *input.Password
// 	}
//
// 	if input.Email != nil {
// 		user.Email = *input.Email
// 	}
//
// 	v := validator.New()
//
// 	if data.ValidateUser(v, user); !v.Valid() {
// 		app.failedValidationResponse(w, r, v.Errors)
// 		return
// 	}
//
// 	err = app.models.Users.Update(user)
//
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 		return
// 	}
//
// 	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
//
// func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := app.readIdParam(r)
// 	if err != nil {
// 		app.notFoundResponse(w, r)
// 		return
// 	}
//
// 	err = app.models.Users.Delete(id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			app.notFoundResponse(w, r)
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}
//
// 	err = app.writeJSON(w, http.StatusOK, envelope{"message": "user was successfully deleted"}, nil)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
