package repository

import (
	"context"

	"meh/core/user"
	"meh/ent"
	eu "meh/ent/user"

	"github.com/pkg/errors"
)

type User struct {
	db *ent.Client
}

func NewUser(db *ent.Client) *User {
	return &User{
		db: db,
	}
}

func (r *User) Create(ctx context.Context, du *user.User) (user.ID, error) {
	res, err := r.db.User.Create().
		SetScreenName(du.ScreenName).
		Save(ctx)

	if err != nil {
		return 0, errors.Wrap(err, "failed to db.User.Create")
	}

	return user.ID(res.ID), nil
}

func (r *User) ExistsByScreenName(ctx context.Context, name string) (bool, error) {
	found, err := r.db.User.Query().Where(
		eu.ScreenNameEQ(name),
	).Exist(ctx)

	if err != nil {
		return false, errors.Wrap(err, "failed to check screenname exsistence")
	}

	return found, nil
}

func (r *User) FindByUserIDs(ctx context.Context, userIDs user.IDs) ([]*user.User, error) {
	es, err := r.db.User.Query().Where(
		eu.IDIn(userIDs.ToUint64s()...),
	).All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "failed to check screenname exsistence")
	}

	var us []*user.User
	for _, e := range es {
		us = append(us, &user.User{
			ID:         user.ID(e.ID),
			ScreenName: e.ScreenName,
		})
	}

	return us, nil
}
