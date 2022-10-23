package data

import "gitlab.com/distributed_lab/kit/pgdb"

type CategoriesQ interface {
	New() CategoriesQ

	Get() (*Category, error)
	Select() ([]Category, error)

	Transaction(fn func(q CategoriesQ) error) error

	Insert(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) CategoriesQ

	FilterById(ids ...int64) CategoriesQ
	FilterByNames(names ...string) CategoriesQ
	FilterByUnits(units ...string) CategoriesQ
}

type Category struct {
	Id           int64  `db:"id" structs:"-"`
	CategoryName string `db:"category_name" structs:"category_name"`
	Unit         string `db:"unit" structs:"unit"`
}
