package meh

import "meh/core/user"

type Timeline struct {
	UserID user.ID
	Mehs   []*Meh
}
