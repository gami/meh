package controller

import (
	"context"

	"meh/core"
	"meh/core/meh"
	"meh/core/user"
	"meh/usecase/form"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, input form.CreateUser) (user.ID, error)
	FindByUserIDs(ctx context.Context, userIDs user.IDs) ([]*user.User, error)
}

type FollowUsecase interface {
	Follow(ctx context.Context, from, to user.ID) error
	Remove(ctx context.Context, from, to user.ID) error
}

type MehUsecase interface {
	Meh(ctx context.Context, input form.CreateMeh) (meh.ID, error)
	ListMehsInTimeline(ctx context.Context, userID user.ID, pagination core.Pagination) (meh.Mehs, core.Pagination, error)
}
