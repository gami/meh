package meh

import (
	"context"

	"meh/core"
	"meh/core/user"
)

type Repository interface {
	Create(ctx context.Context, meh *Meh) (ID, error)
	AddToTimeline(ctx context.Context, ID ID, followeeID user.ID) error
	ListMehsInTimeline(ctx context.Context, userID user.ID, pagination core.Pagination) (Mehs, core.Pagination, error)
}
