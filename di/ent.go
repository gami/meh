package di

import (
	"meh/ent"
	"meh/mysql"
)

func InjectEntClient() *ent.Client {
	return mysql.NewEntClient()
}
