package receipt

type Receipt struct {
	Id           int
	MealId       int `db:"meal"`
	IngredientId int `db:"ingredient"`
	Quantity     int
}

func NewReceipt(ingredientId int, quantity int) *Receipt {
	return &Receipt{IngredientId: ingredientId, Quantity: quantity}
}
