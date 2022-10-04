package meal

type Meal struct {
	Id         int
	MealName   string `db:"position_name"`
	CategoryId int    `db:"category"`
	Price      float64
	Amount     float64
	MenuId     int `db:"menu"`
}

func NewMeal(mealName string, categoryId int, price float64, amount float64, menuId int) *Meal {
	return &Meal{MealName: mealName, CategoryId: categoryId, Price: price, Amount: amount, MenuId: menuId}
}
