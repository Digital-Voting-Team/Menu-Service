package meal

type Meal struct {
	Id         int
	MealName   string `db:"position_name"`
	CategoryId int    `db:"category"`
	Price      float64
	Amount     float64
}

func NewMeal(mealName string, categoryId int, price float64, amount float64) *Meal {
	return &Meal{MealName: mealName, CategoryId: categoryId, Price: price, Amount: amount}
}
