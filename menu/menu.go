package menu

type Menu struct {
	Id     int
	CafeId int `db:"cafe_id"`
}

func NewMenu(cafeId int) *Menu {
	return &Menu{CafeId: cafeId}
}
