package data

import "gitlab.com/distributed_lab/kit/pgdb"

type ReceiptsQ interface {
	New() ReceiptsQ

	Get() (*Receipt, error)
	Select() ([]Receipt, error)

	Transaction(fn func(q ReceiptsQ) error) error

	Insert(receipt Receipt) (Receipt, error)
	Update(receipt Receipt) (Receipt, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) ReceiptsQ

	FilterById(ids ...int64) ReceiptsQ
	FilterByMealId(ids ...int64) ReceiptsQ
	FilterByIngredientId(ids ...int64) ReceiptsQ
	FilterByQuantityFrom(quantities ...int64) ReceiptsQ
	FilterByQuantityTo(quantities ...int64) ReceiptsQ

	JoinMeal() ReceiptsQ
}

type Receipt struct {
	Id           int64 `db:"id" structs:"-"`
	MealId       int64 `db:"meal_id" structs:"meal_id"`
	IngredientId int64 `db:"ingredient_id" structs:"ingredient_id"`
	Quantity     int64 `db:"quantity" structs:"quantity"`
}
