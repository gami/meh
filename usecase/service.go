package usecase

import (
	"context"

	"meh/core"
	"meh/core/meh"
	"meh/core/user"
)

type UserService interface {
	Create(ctx context.Context, us *user.User) (user.ID, error)
	FindByUserIDs(ctx context.Context, userIDs user.IDs) ([]*user.User, error)
}

type FollowService interface {
	Follow(ctx context.Context, from, to user.ID) error
	Remove(ctx context.Context, from, to user.ID) error
	ListFollowers(ctx context.Context, userID user.ID) ([]user.ID, error)
}

type MehService interface {
	Create(ctx context.Context, meh *meh.Meh) (meh.ID, error)
	AddToTimeline(ctx context.Context, id meh.ID, followeeIDs []user.ID) error
	ListMehsInTimeline(ctx context.Context, userID user.ID, pagination core.Pagination) (meh.Mehs, core.Pagination, error)
}
