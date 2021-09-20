package usecase

import (
	"context"
	"fmt"
	"log"

	"meh/core"
	"meh/core/meh"
	"meh/core/user"
	"meh/usecase/form"

	"github.com/pkg/errors"
)

type Meh struct {
	tx     core.Tx
	meh    MehService
	follow FollowService
}

func NewMeh(tx core.Tx, mh MehService, fl FollowService) *Meh {
	return &Meh{
		tx:     tx,
		meh:    mh,
		follow: fl,
	}
}

func (u *Meh) Meh(ctx context.Context, input form.CreateMeh) (meh.ID, error) {
	if err := input.Validate(); err != nil {
		return 0, errors.Wrap(err, "failed to validate")
	}

	res, err := u.tx.Transact(ctx, func(ctx context.Context) (interface{}, error) {
		id, err := u.meh.Create(ctx, &meh.Meh{
			UserID: input.UserID,
			Text:   input.Text,
		})

		if err != nil {
			return 0, errors.Wrap(err, "failed to create meh")
		}

		return id, nil
	})

	if err != nil {
		return 0, err
	}

	id, ok := res.(meh.ID)
	if !ok {
		return 0, fmt.Errorf("type must be meh.ID but `%T`", res)
	}

	go func() {
		ctx = context.Background()
		targets, err := u.follow.ListFollowers(ctx, input.UserID)
		if err != nil {
			log.Println(errors.Wrap(err, "failed to create meh"))
			return
		}

		targets = append(targets, input.UserID) // 自分のタイムラインにも表示される

		// なんらかの原因でタイムラインに追加できない投稿があった場合も、
		// 追加済みのタイムラインや投稿を無効にする必要はなく、
		// むしろパフォーマンスが重視されるので、ここでのTransactionは不要
		if err := u.meh.AddToTimeline(ctx, id, targets); err != nil {
			log.Println(errors.Wrap(err, "failed to create meh"))
		}
	}()

	return id, nil
}

func (u *Meh) ListMehsInTimeline(ctx context.Context, userID user.ID, pagination core.Pagination) (meh.Mehs, core.Pagination, error) {
	return u.meh.ListMehsInTimeline(ctx, userID, pagination)
}
