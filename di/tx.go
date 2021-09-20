package di

import (
	"meh/core"
	"meh/mysql"
)

func InjectTx() core.Tx {
	return mysql.NewDBTx(
		InjectEntClient(),
	)
}
