package graph

import (
	"database/sql"

	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func Connect() *bun.DB {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	connStr := viper.GetString("DB_URL")

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connStr)))
	bunDb := bun.NewDB(sqldb, pgdialect.New())
	return bunDb
}
