package database

import (
	u "awesomeProject/utils"
	"context"
	"database/sql"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var JetDB *sql.DB
var Ctx = context.TODO()

//var NotFoundError = errors.New("not found")

func ConnectDatabase() {
	db, err := sql.Open("postgres", u.Settings.Database)
	if err != nil {
		panic("failed to connect database")
	}
	JetDB = db

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}

func InInt(ids []int64) []Expression {
	var ids2 []Expression
	for _, id := range ids {
		ids2 = append(ids2, Int(id))
	}
	return ids2
}

func InString(ids []string) []Expression {
	var ids2 []Expression
	for _, id := range ids {
		ids2 = append(ids2, String(id))
	}
	return ids2
}

func NotDeleted(DeletedAt ColumnTimestampz) BoolExpression {
	return DeletedAt.GT(RawTimestampz("now()"))
}
