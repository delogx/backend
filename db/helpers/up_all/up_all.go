package main

import (
	"backend/db"
	"backend/src/initializers"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	initializers.LoadEnv()
	db.ConnectDB()
	db, err := db.DB.DB()
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	for {
		err = m.Up()
		if err != nil {
			if err == migrate.ErrNoChange {
				break
			}
			panic(err)
		}
	}
}
