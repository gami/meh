package repository

import (
	"context"
	"math"

	"meh/core"
	"meh/core/meh"
	"meh/core/user"
	"meh/ent"
	"meh/ent/timeline"
	"meh/pkg/numbers"

	"github.com/pkg/errors"
)

type Meh struct {
	db *ent.Client
}

func NewMeh(db *ent.Client) *Meh {
	return &Meh{
		db: db,
	}
}

func (r *Meh) Create(ctx context.Context, mh *meh.Meh) (meh.ID, error) {
	em, err := r.db.Meh.Create().
		SetUserID(uint64(mh.UserID)).
		SetText(string(mh.Text)).
		Save(ctx)

	if err != nil {
		return 0, errors.Wrap(err, "failed to create meh")
	}

	return meh.ID(em.ID), nil
}

func (r *Meh) AddToTimeline(ctx context.Context, mehID meh.ID, followeeID user.ID) error {
	_, err := r.db.Timeline.Create().
		SetMehID(uint64(mehID)).
		SetUserID(uint64(followeeID)).
		Save(ctx)

	if err != nil {
		return errors.Wrap(err, "failed to add timeline")
	}

	return nil
}

func (r *Meh) ListMehsInTimeline(ctx context.Context, userID user.ID, pg core.Pagination) (meh.Mehs, core.Pagination, error) {
	var lastID uint64
	if pg.LastID == nil {
		lastID = math.MaxUint64
	} else {
		lastID = *pg.LastID
	}

	ts, err := r.db.Timeline.Query().
		Where(
			timeline.UserIDEQ(uint64(userID)),
			timeline.IDLT(lastID),
		).
		Limit(int(pg.Count)).
		Order(ent.Desc("id")).
		WithMeh().All(ctx)
	if err != nil {
		return nil, core.Pagination{}, errors.Wrap(err, "failed to query timeline")
	}

	if len(ts) == 0 {
		pg.LastID = numbers.PointerUint64(0)
		return nil, pg, nil
	}

	var mhs meh.Mehs
	for _, t := range ts {
		m := t.Edges.Meh

		mhs = append(mhs, &meh.Meh{
			ID:        meh.ID(m.ID),
			UserID:    user.ID(m.UserID),
			Text:      meh.Text(m.Text),
			CreatedAt: m.CreatedAt,
		})
	}

	pg.LastID = numbers.PointerUint64(ts[len(ts)-1].ID)

	return mhs, pg, nil
}
