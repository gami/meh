package core

import (
	"context"
)

type Tx interface {
	Transact(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
