package usecase

import (
	"context"

	"meh/core"
	"meh/core/user"
	"meh/usecase/form"

	"github.com/pkg/errors"
)

type User struct {
	user UserService
	tx   core.Tx
}

func NewUser(tx core.Tx, us UserService) *User {
	return &User{
		user: us,
		tx:   tx,
	}
}

func (u *User) CreateUser(ctx context.Context, input form.CreateUser) (user.ID, error) {
	if err := input.Validate(); err != nil {
		return 0, errors.Wrap(err, "failed to validate create user")
	}

	res, err := u.tx.Transact(ctx, func(ctx context.Context) (interface{}, error) {
		id, err := u.user.Create(ctx, &user.User{
			ScreenName: input.ScreenName,
		})
		if err != nil {
			return 0, errors.Wrap(err, "failed to create user")
		}

		return id, nil
	})

	if err != nil {
		return 0, err
	}

	return res.(user.ID), nil
}

func (u *User) FindByUserIDs(ctx context.Context, userIDs user.IDs) ([]*user.User, error) {
	return u.user.FindByUserIDs(ctx, userIDs)
}
