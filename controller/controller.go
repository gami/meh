package controller

import (
	"encoding/json"
	"log"
	"net/http"

	openapi "meh/api/openapi"

	"github.com/pkg/errors"
)

type Controller struct {
	*User
	*Follow
	*Meh
}

func NewController(
	u *User,
	f *Follow,
	m *Meh,
) *Controller {
	return &Controller{
		User:   u,
		Follow: f,
		Meh:    m,
	}
}

func RespondOK(w http.ResponseWriter, result interface{}) {
	w.WriteHeader(http.StatusOK)

	if result == nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		RespondError(w, errors.Wrap(err, "faile to respond_ok"))

		return
	}
}

func RespondError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(err)

	w.WriteHeader(http.StatusInternalServerError) // TODO カスタムエラーを入れてステータス別に処理を分ける
	_ = json.NewEncoder(w).Encode(&openapi.Error{
		Message: err.Error(),
	})
}
