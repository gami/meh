package form

import (
	"errors"

	"meh/core/meh"
	"meh/core/user"
)

type CreateMeh struct {
	UserID user.ID
	Text   meh.Text
}

func (f *CreateMeh) Validate() error {
	if f.UserID < 1 {
		return errors.New("userID is empty")
	}

	return f.Text.Validate()
}
