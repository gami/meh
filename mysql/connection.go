package mysql

import (
	"fmt"
	"log"
	"time"

	"entgo.io/ent/dialect/sql"

	"meh/config"
	"meh/ent"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	entClient *ent.Client
	loc       = time.FixedZone("Asia/Tokyo", 9*60*60)
)

const (
	ConnMaxLifetimeSec = 120
	MaxOpenConns       = 100
	MaxIdleConns       = 100
)

func init() {
	err := connectDB()
	if err != nil {
		log.Println(err)
		return
	}
}

func connectWithConfig(connStr string) (*ent.Client, error) {
	drv, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	db := drv.DB()

	db.SetConnMaxLifetime(ConnMaxLifetimeSec * time.Second)
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)

	return ent.NewClient(ent.Driver(drv)), nil
}

func dsn(database config.Database) string {
	mc := mysql.Config{
		User:   database.User,
		Passwd: database.Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", database.Host, database.Port),
		DBName: database.Name,
		Params: map[string]string{
			"sql_mode": "TRADITIONAL",
		},
		Collation:            "utf8mb4_bin",
		Loc:                  loc,
		CheckConnLiveness:    true,
		ParseTime:            true,
		MaxAllowedPacket:     1024 * 1024 * 16,
		AllowNativePasswords: true,
	}

	return mc.FormatDSN()
}

func connectDB() error {
	c, err := connectWithConfig(dsn(config.GetConfig().DB))
	if err != nil {
		return errors.Wrap(err, "failed to connect db")
	}

	entClient = c
	return nil
}

func NewEntClient() *ent.Client {
	return entClient
}
