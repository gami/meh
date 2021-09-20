package repository

import (
	"context"

	"meh/core/user"
	"meh/ent"
	"meh/ent/follow"

	"github.com/pkg/errors"
)

type Follow struct {
	db *ent.Client
}

func NewFollow(db *ent.Client) *Follow {
	return &Follow{
		db: db,
	}
}

func (r *Follow) Create(ctx context.Context, from, to user.ID) error {
	_, err := r.db.Follow.Create().
		SetUserID(uint64(from)).
		SetFolloweeID(uint64(to)).Save(ctx)

	if err != nil {
		return errors.Wrap(err, "failed to create follow")
	}

	return nil
}

func (r *Follow) Delete(ctx context.Context, from, to user.ID) error {
	cnt, err := r.db.Follow.Delete().Where(
		follow.UserIDEQ(uint64(from)),
		follow.FolloweeIDEQ(uint64(to)),
	).Exec(ctx)

	if err != nil {
		return errors.Wrap(err, "failed to delete follow")
	}

	if cnt == 0 {
		return errors.New("no deleted follow")
	}

	return nil
}

func (r *Follow) Exists(ctx context.Context, from, to user.ID) (bool, error) {
	found, err := r.db.Follow.Query().Where(
		follow.UserIDEQ(uint64(from)),
		follow.FolloweeIDEQ(uint64(to)),
	).Exist(ctx)

	if err != nil {
		return false, errors.Wrap(err, "failed to check follow exsitence")
	}

	return found, nil
}

func (r *Follow) ListFollowers(ctx context.Context, userID user.ID) ([]user.ID, error) {
	es, err := r.db.Follow.Query().Where(
		follow.FolloweeIDEQ(uint64(userID)),
	).All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "failed to list followers")
	}

	var fids []user.ID
	for _, e := range es {
		fids = append(fids, user.ID(e.UserID))
	}

	return fids, nil
}
