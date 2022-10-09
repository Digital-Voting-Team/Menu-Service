package main

import (
	"Menu-Service/category"
	"Menu-Service/meal"
	"Menu-Service/menu"
	"Menu-Service/receipt"
	"Menu-Service/utils"
	"github.com/jmoiron/sqlx"
	"log"
)

func CategoriesSimulation(db *sqlx.DB) {
	repo := category.NewRepository(db)
	newEntity := utils.GenerateMockCategory()

	id, err := repo.Insert(newEntity)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id of the added category: %d", id)

	array, err := repo.Select()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\ncategories: %+v", array)
}

func MealsSimulation(db *sqlx.DB) {
	repo := meal.NewRepository(db)
	newEntity := utils.GenerateMockMeal()

	id, err := repo.Insert(newEntity)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id of the added meal: %d", id)

	array, err := repo.Select()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nmeals: %+v", array)
}

func MenusSimulation(db *sqlx.DB) {
	repo := menu.NewRepository(db)
	newEntity := utils.GenerateMockMenu()

	id, err := repo.Insert(newEntity)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id of the added menu: %d", id)

	array, err := repo.Select()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nmenus: %+v", array)
}

func ReceiptsSimulation(db *sqlx.DB) {
	repo := receipt.NewRepository(db)
	newEntity := utils.GenerateMockReceipt()

	id, err := repo.Insert(newEntity)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id of the added receipt: %d", id)

	array, err := repo.Select()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nreceipts: %+v", array)
}
