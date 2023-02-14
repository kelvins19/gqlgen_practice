package graph

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() *bun.DB {
	connStr := os.Getenv("DB_URL")
	// opt, err := pg.ParseURL(connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// db := pg.Connect(opt)
	// return db

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))
	bunDb := bun.NewDB(sqldb, pgdialect.New())
	return bunDb
}
