package utils

import (
	"Menu-Service/category"
	"Menu-Service/meal"
	"Menu-Service/menu"
	"Menu-Service/receipt"
	"math/rand"
)

func GenerateMockCategory() *category.Category {
	return &category.Category{
		CategoryName: RandStringRunes(10),
		Unit:         "ml",
	}
}

func GenerateMockMeal() *meal.Meal {
	return &meal.Meal{
		MealName:   RandStringRunes(10),
		CategoryId: 1,
		MenuId:     1,
	}
}

func GenerateMockMenu() *menu.Menu {
	return &menu.Menu{
		CafeId: 1,
	}
}

func GenerateMockReceipt() *receipt.Receipt {
	return &receipt.Receipt{
		MealId:       1,
		IngredientId: 1,
		Quantity:     rand.Intn(20),
	}
}
