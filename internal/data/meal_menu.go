package data

import "gitlab.com/distributed_lab/kit/pgdb"

type MealMenusQ interface {
	New() MealMenusQ

	Get() (*MealMenu, error)
	Select() ([]MealMenu, error)

	Transaction(fn func(q MealMenusQ) error) error

	Insert(address MealMenu) (MealMenu, error)
	Update(address MealMenu) (MealMenu, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) MealMenusQ

	FilterById(ids ...int64) MealMenusQ
	FilterByMealId(ids ...int64) MealMenusQ
	FilterByMenuId(ids ...int64) MealMenusQ

	JoinMeal() MealMenusQ
	JoinMenu() MealMenusQ
}

type MealMenu struct {
	Id     int64 `db:"id" structs:"-"`
	MealId int64 `db:"meal_id" structs:"meal_id"`
	MenuId int64 `db:"menu_id" structs:"menu_id"`
}
