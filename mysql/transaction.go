package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"meh/ent"

	"github.com/pkg/errors"
)

type txKey struct{}

var ctxKey txKey

type DBTx struct {
	db *ent.Client
}

func NewDBTx(db *ent.Client) *DBTx {
	return &DBTx{
		db: db,
	}
}

func (t *DBTx) Transact(ctx context.Context, process func(context.Context) (interface{}, error)) (interface{}, error) {
	if InTransaction(ctx) {
		log.Println("this context has already other transaction")

		return process(ctx)
	}

	defer t.clear(ctx)

	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction failed")
	}

	ctx = context.WithValue(ctx, ctxKey, tx)

	obj, err := process(ctx)
	if err == nil {
		err = tx.Commit()
	}

	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			err = errors.Wrap(err, fmt.Sprintf("rollback err = %v", errR))
		}

		return nil, err
	}

	return obj, nil
}

func (t *DBTx) clear(ctx context.Context) {
	_ = context.WithValue(ctx, ctxKey, nil)
}

func GetTx(ctx context.Context) (*ent.Tx, bool) {
	v := ctx.Value(ctxKey)

	tx, ok := v.(*ent.Tx)
	if !ok {
		return nil, false
	}

	return tx, true
}

func InTransaction(ctx context.Context) bool {
	_, already := GetTx(ctx)

	return already
}
