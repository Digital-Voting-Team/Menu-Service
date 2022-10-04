package category

type Category struct {
	Id           int
	CategoryName string `db:"category_name"`
	Unit         string
}

func NewCategory(categoryName string, unit string) *Category {
	return &Category{CategoryName: categoryName, Unit: unit}
}
