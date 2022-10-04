package receipt

type Receipt struct {
	Id           int
	MealId       int `db:"meal_id"`
	IngredientId int `db:"ingredient_id"`
	Quantity     int
}

func NewReceipt(ingredientId int, quantity int) *Receipt {
	return &Receipt{IngredientId: ingredientId, Quantity: quantity}
}
