package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/meal_menu"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateMealMenu(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateMealMenuRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	MealMenu := data.MealMenu{
		MealId: cast.ToInt64(request.Data.Relationships.Meal.Data.ID),
		MenuId: cast.ToInt64(request.Data.Relationships.Menu.Data.ID),
	}

	var resultMealMenu data.MealMenu
	relateMeal, err := helpers.MealsQ(r).FilterById(MealMenu.MealId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get meal")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	relateMenu, err := helpers.MenusQ(r).FilterById(MealMenu.MenuId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get menu")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultMealMenu, err = helpers.MealMenusQ(r).Insert(MealMenu)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create mealMenu")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Meal{
		Key: resources.NewKeyInt64(relateMeal.Id, resources.MEAL),
		Attributes: resources.MealAttributes{
			MealName: relateMeal.MealName,
			Price:    relateMeal.Price,
			Amount:   relateMeal.Amount,
		},
		Relationships: resources.MealRelationships{
			Category: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(relateMeal.CategoryId, 10),
					Type: resources.CATEGORY,
				},
			},
		},
	})

	includes.Add(&resources.Menu{
		Key: resources.NewKeyInt64(relateMenu.Id, resources.MENU),
		Relationships: resources.MenuRelationships{
			Cafe: resources.Relation{
				Data: &resources.Key{
					ID:   strconv.FormatInt(relateMenu.CafeId, 10),
					Type: resources.CAFE_REF,
				},
			},
		},
	})

	result := resources.MealMenuResponse{
		Data: resources.MealMenu{
			Key: resources.NewKeyInt64(resultMealMenu.Id, resources.MEAL_MENU),
			Relationships: resources.MealMenuRelationships{
				Meal: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultMealMenu.MealId, 10),
						Type: resources.MEAL,
					},
				},
				Menu: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultMealMenu.MenuId, 10),
						Type: resources.MENU,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
