package follow

import (
	"context"

	"meh/core/user"
)

type Repository interface {
	Create(ctx context.Context, from user.ID, to user.ID) error
	Delete(ctx context.Context, from user.ID, to user.ID) error
	Exists(ctx context.Context, from user.ID, to user.ID) (bool, error)
	ListFollowers(ctx context.Context, userID user.ID) ([]user.ID, error)
}
