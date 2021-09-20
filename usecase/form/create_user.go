package form

import "errors"

type CreateUser struct {
	ScreenName string
}

func (f *CreateUser) Validate() error {
	if f.ScreenName == "" {
		return errors.New("screen name is empty")
	}

	return nil
}
