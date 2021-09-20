package controller

import (
	"encoding/json"
	"net/http"

	openapi "meh/api/openapi"
	"meh/core/user"
)

type Follow struct {
	follow FollowUsecase
}

func NewFollow(fl FollowUsecase) *Follow {
	return &Follow{
		follow: fl,
	}
}

// (POST /follows/create)
func (c *Follow) FollowUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.FollowUserJSONBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, err)

		return
	}

	// TODO body.UserIDとctxの中のUserIDが一致するか検証する
	// 認可をどのレイヤーで行うかは難しいが、このアプリケーションであれば、
	// そのリソースのオーナーであるかどうか（もしくは管理者/Systemかどうか）の判定しかなさそうなので、
	// Controllerで良さそう
	// ユーザーにロールがあって、ロールによってできることを細かく制御する必要がある場合は
	// Usecaseで行って、ロジックはcoreにおきたい

	err := c.follow.Follow(ctx, user.ID(body.UserId), user.ID(body.FolloweeId))
	if err != nil {
		RespondError(w, err)

		return
	}

	RespondOK(w, nil)
}

// (POST /follows/delete)
func (c *Follow) RemoveUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.FollowUserJSONBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, err)

		return
	}

	err := c.follow.Remove(ctx, user.ID(body.UserId), user.ID(body.FolloweeId))
	if err != nil {
		RespondError(w, err)

		return
	}

	RespondOK(w, nil)
}
