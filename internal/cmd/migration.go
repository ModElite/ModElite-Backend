package main

import (
	"log"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	configEnv, err := config.GetEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize migrator
	m, err := migrate.New(
		"file://db/migrations",
		configEnv.DATABASE_URI,
	)
	if err != nil {
		log.Fatal("Cannot run migration", err)
	}

	selectCase := 1

	if selectCase == 1 {
		err = m.Up()
		if err != nil {
			log.Fatal("Cannot run migration Up", err)
		}
		log.Default().Println("Migration Up done")
	} else if selectCase == 2 {
		err = m.Down()
		if err != nil {
			log.Fatal("Cannot run migration Down", err)
		}
		log.Default().Println("Migration Down done")
	} else {
		err = m.Drop()
		if err != nil {
			log.Fatal("Cannot run migration Drop", err)
		}
		log.Default().Println("Migration Drop done")
	}
}
