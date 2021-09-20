package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	openapi "meh/api/openapi"
	"meh/core"
	"meh/core/meh"
	"meh/core/user"
	"meh/usecase/form"
)

type Meh struct {
	meh  MehUsecase
	user UserUsecase
}

func NewMeh(mh MehUsecase, us UserUsecase) *Meh {
	return &Meh{
		meh:  mh,
		user: us,
	}
}

// (POST /mehs)
func (c *Meh) CreateMeh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.CreateMehJSONBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, err)

		return
	}

	_, err := c.meh.Meh(ctx, form.CreateMeh{
		UserID: user.ID(body.UserId),
		Text:   meh.Text(body.Text),
	})
	if err != nil {
		RespondError(w, err)

		return
	}

	RespondOK(w, nil)
}

// (GET /me/timeline)
func (c *Meh) ShowTimeline(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body openapi.ShowTimelineJSONBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondError(w, err)

		return
	}

	userIDValue := ctx.Value(user.IDContextKey)
	userID, ok := userIDValue.(uint64)
	if !ok {
		RespondError(w, errors.New("missing userID"))

		return
	}

	pg := core.Pagination{
		LastID: body.Pagination.LastId,
		Count:  body.Pagination.Count,
	}

	mehs, pg, err := c.meh.ListMehsInTimeline(ctx, user.ID(userID), pg)
	if err != nil {
		RespondError(w, err)

		return
	}

	users, err := c.user.FindByUserIDs(ctx, mehs.UserIDs())
	if err != nil {
		RespondError(w, err)

		return
	}

	var res []openapi.Meh
	for _, m := range mehs {
		var mu *openapi.User
		for _, u := range users {
			if u.ID == m.UserID {
				mu = &openapi.User{
					Id:         uint64(u.ID),
					ScreenName: u.ScreenName,
				}
			}
		}

		if mu == nil {
			RespondError(w, fmt.Errorf("missing user for meh (%d)", m.ID))

			return
		}

		res = append(res, openapi.Meh{
			Id:   (*uint64)(&m.ID),
			Text: string(m.Text),
			User: mu,
		})
	}

	RespondOK(w, &openapi.TimelineResponse{
		Mehs: res,
		Pagination: openapi.Pagination{
			Count:  pg.Count,
			LastId: pg.LastID,
		},
	})
}
