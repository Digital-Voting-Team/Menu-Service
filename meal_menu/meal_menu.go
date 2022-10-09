package meal_menu

type MealMenu struct {
	Id     int
	MealId int `db:"meal"`
	MenuId int `db:"menu"`
}
