package data

import "gitlab.com/distributed_lab/kit/pgdb"

type MealsQ interface {
	New() MealsQ

	Get() (*Meal, error)
	Select() ([]Meal, error)

	Transaction(fn func(q MealsQ) error) error

	Insert(address Meal) (Meal, error)
	Update(address Meal) (Meal, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) MealsQ

	FilterById(ids ...int64) MealsQ
	FilterByNames(names ...string) MealsQ
	FilterByPriceFrom(prices ...float64) MealsQ
	FilterByPriceTo(prices ...float64) MealsQ
	FilterByAmount(amounts ...float64) MealsQ
	FilterByCategoryId(ids ...float64) MealsQ

	JoinCategory() MealsQ
}

type Meal struct {
	Id         int64   `db:"id" structs:"-"`
	MealName   string  `db:"meal_name" structs:"meal_name"`
	CategoryId int64   `db:"category_id" structs:"category_id"`
	Price      float64 `db:"price" structs:"price"`
	Amount     float64 `db:"amount" structs:"amount"`
}
