package main

import (
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/i4ox/golang-project-template/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	dbUrl := cfg.Database.URL

	// Get all flags
	action := flag.String("action", "up", "migration action (up, down, force)")
	steps := flag.Int("steps", 1, "number of steps for 'up' or 'down' actions")
	flag.Parse()

	m, err := migrate.New("file://./migrations", dbUrl)
	if err != nil {
		log.Fatalf("Could not create migration instance: %v", err)
	}

	switch *action {
	case "up":
		if *steps > 1 {
			err = m.Steps(*steps)
		} else {
			err = m.Up()
		}
	case "down":
		if *steps > 1 {
			err = m.Steps(-(*steps))
		} else {
			err = m.Down()
		}
	case "force":
		err = m.Force(*steps)
	default:
		log.Fatalf("Unknown action %s", *action)
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration action completed:", *action)
}
