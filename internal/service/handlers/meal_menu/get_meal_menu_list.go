package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/meal_menu"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetMealMenuList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetMealMenuListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	mealMenusQ := helpers.MealMenusQ(r)
	applyFilters(mealMenusQ, request)
	mealMenus, err := mealMenusQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get mealMenus")
		ape.Render(w, problems.InternalError())
		return
	}
	meals, err := helpers.MealsQ(r).FilterById(getMealIds(mealMenus)...).Select()
	menus, err := helpers.MenusQ(r).FilterById(getMenuIds(mealMenus)...).Select()

	response := resources.MealMenuListResponse{
		Data:     newMealMenusList(mealMenus),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newMealMenuIncluded(meals, menus),
	}
	ape.Render(w, response)
}

func applyFilters(q data.MealMenusQ, request requests.GetMealMenuListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterMealId) > 0 {
		q.FilterByMealId(request.FilterMealId...)
	}

	if len(request.FilterMenuId) > 0 {
		q.FilterByMealId(request.FilterMenuId...)
	}
}

func newMealMenusList(mealMenus []data.MealMenu) []resources.MealMenu {
	result := make([]resources.MealMenu, len(mealMenus))
	for i, mealMenu := range mealMenus {
		result[i] = resources.MealMenu{
			Key: resources.NewKeyInt64(mealMenu.Id, resources.MEAL_MENU),
			Relationships: resources.MealMenuRelationships{
				Meal: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(mealMenu.MealId, 10),
						Type: resources.MEAL,
					},
				},
				Menu: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(mealMenu.MenuId, 10),
						Type: resources.MENU,
					},
				},
			},
		}
	}
	return result
}

func getMealIds(mealMenus []data.MealMenu) []int64 {
	mealIDs := make([]int64, len(mealMenus))
	for i := 0; i < len(mealMenus); i++ {
		mealIDs[i] = mealMenus[i].MealId
	}
	return mealIDs
}

func getMenuIds(mealMenus []data.MealMenu) []int64 {
	menuIDs := make([]int64, len(mealMenus))
	for i := 0; i < len(mealMenus); i++ {
		menuIDs[i] = mealMenus[i].MenuId
	}
	return menuIDs
}

func newMealMenuIncluded(meals []data.Meal, menus []data.Menu) resources.Included {
	result := resources.Included{}
	for _, item := range meals {
		resource := newMealModel(item)
		result.Add(&resource)
	}
	for _, item := range menus {
		resource := newMenuModel(item)
		result.Add(&resource)
	}
	return result
}

func newMealModel(meal data.Meal) resources.Meal {
	return resources.Meal{
		Key: resources.NewKeyInt64(meal.Id, resources.MEAL),
		Attributes: resources.MealAttributes{
			MealName: meal.MealName,
			Price:    meal.Price,
			Amount:   meal.Amount,
		},
		Relationships: resources.MealRelationships{
			Category: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(meal.CategoryId, 10),
					Type: resources.CATEGORY,
				},
			},
		},
	}
}

func newMenuModel(menu data.Menu) resources.Menu {
	return resources.Menu{
		Key: resources.NewKeyInt64(menu.Id, resources.MENU),
		Relationships: resources.MenuRelationships{
			Cafe: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(menu.CafeId, 10),
					Type: resources.CAFE_REF,
				},
			},
		},
	}
}
