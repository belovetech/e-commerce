package seeders

import (
	"context"
	"fmt"
	"log"

	"github.com/belovetech/e-commerce/database/sqlc"
)

func getSeeders() []Seeder {
	return []Seeder{
		AdminSeeder{},
		ProductSeeder{},
	}
}

func RunSeeders(queries *sqlc.Queries) error {
	for _, seeder := range getSeeders() {
		executed, err := isSeederExecuted(queries, seeder.Name())
		if err != nil {
			fmt.Println("err: ", err)
			return err
		}

		if executed {
			log.Printf("Seeder '%s' has already run. Skipping...\n", seeder.Name())
			continue
		}

		log.Printf("Running seeder '%s'...\n", seeder.Name())
		if err := seeder.Seed(queries); err != nil {
			return err
		}

		if err := recordSeederExecution(queries, seeder.Name()); err != nil {
			return err
		}
	}
	return nil
}

func isSeederExecuted(queries *sqlc.Queries, seederName string) (bool, error) {
	history, _ := queries.GetSeederByName(context.Background(), seederName)
	return history != "", nil
}

func recordSeederExecution(queries *sqlc.Queries, seederName string) error {
	return queries.CreateSeederHistory(context.Background(), seederName)
}
