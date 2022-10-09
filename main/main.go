package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func Connect(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connStr)
	return db, err
}

func main() {
	connStr := "user=postgres dbname=Menu sslmode=disable password=password"
	db, err := Connect(connStr)

	if err != nil {
		log.Fatal(err)
	}

	CategoriesSimulation(db)
	MenusSimulation(db)
	MealsSimulation(db)
	ReceiptsSimulation(db)
	MealMenuSimulation(db)

	Clean(db)
}
