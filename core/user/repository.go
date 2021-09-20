package user

import (
	"context"
)

type Repository interface {
	ExistsByScreenName(ctx context.Context, name string) (bool, error)
	FindByUserIDs(ctx context.Context, userIDs IDs) ([]*User, error)
	Create(ctx context.Context, user *User) (ID, error)
}
