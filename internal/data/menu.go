package data

import "gitlab.com/distributed_lab/kit/pgdb"

type MenusQ interface {
	New() MenusQ

	Get() (*Menu, error)
	Select() ([]Menu, error)

	Transaction(fn func(q MenusQ) error) error

	Insert(address Menu) (Menu, error)
	Update(address Menu) (Menu, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) MenusQ

	FilterById(ids ...int64) MenusQ
	FilterByCafeId(ids ...int64) MenusQ
}

type Menu struct {
	Id     int64 `db:"id" structs:"-"`
	CafeId int64 `db:"cafe_id" structs:"cafe_id"`
}
