package meh

import (
	"errors"
	"time"

	"meh/core/user"
)

const (
	textMaxLength = 140
)

type ID uint64
type Text string
type Mehs []*Meh

type Meh struct {
	ID        ID
	UserID    user.ID
	Text      Text
	CreatedAt time.Time
}

func (t Text) Validate() error {
	if len(t) == 0 {
		return errors.New("text is empty")
	}

	if len([]rune(t)) > textMaxLength {
		return errors.New("text is too long")
	}

	return nil
}

func (ms Mehs) IDs() []ID {
	var ids []ID
	for _, m := range ms {
		ids = append(ids, m.ID)
	}

	return ids
}

func (ms Mehs) UserIDs() []user.ID {
	var ids []user.ID
	for _, m := range ms {
		ids = append(ids, m.UserID)
	}

	return ids
}
