package controller

import (
	"encoding/json"
	"net/http"

	openapi "meh/api/openapi"
	"meh/core/user"
	"meh/usecase/form"
)

type User struct {
	user UserUsecase
}

func NewUser(us UserUsecase) *User {
	return &User{
		user: us,
	}
}

// (POST /users)
func (c *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.CreateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, err)

		return
	}

	id, err := c.user.CreateUser(ctx, form.CreateUser{
		ScreenName: body.ScreenName,
	})
	if err != nil {
		RespondError(w, err)

		return
	}

	RespondOK(w, &user.User{
		ID:         id,
		ScreenName: body.ScreenName,
	})
}
