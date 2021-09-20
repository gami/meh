package usecase

import (
	"context"

	"meh/core"
	"meh/core/user"

	"github.com/pkg/errors"
)

type Follow struct {
	tx     core.Tx
	follow FollowService
}

func NewFollow(tx core.Tx, fl FollowService) *Follow {
	return &Follow{
		tx:     tx,
		follow: fl,
	}
}

func (u *Follow) Follow(ctx context.Context, from, to user.ID) error {
	_, err := u.tx.Transact(ctx, func(ctx context.Context) (interface{}, error) {
		if err := u.follow.Follow(ctx, from, to); err != nil {
			return nil, errors.Wrapf(err, "failed to follow from=%v to=%v", from, to)
		}

		return nil, nil
	})

	if err != nil {
		return err
	}

	return err
}

func (u *Follow) Remove(ctx context.Context, from, to user.ID) error {
	_, err := u.tx.Transact(ctx, func(ctx context.Context) (interface{}, error) {
		if err := u.follow.Remove(ctx, from, to); err != nil {
			return nil, errors.Wrapf(err, "failed to follow from=%v to=%v", from, to)
		}

		return nil, nil
	})

	return err
}
