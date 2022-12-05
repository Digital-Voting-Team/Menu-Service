package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/meal"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateMeal(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateMealRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Meal := data.Meal{
		MealName:   request.Data.Attributes.MealName,
		Price:      request.Data.Attributes.Price,
		Amount:     request.Data.Attributes.Amount,
		CategoryId: cast.ToInt64(request.Data.Relationships.Category.Data.ID),
	}

	var resultMeal data.Meal
	relateCategory, err := helpers.CategoriesQ(r).FilterById(Meal.CategoryId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get category")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultMeal, err = helpers.MealsQ(r).Insert(Meal)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create meal")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var includes resources.Included
	includes.Add(&resources.Category{
		Key: resources.NewKeyInt64(relateCategory.Id, resources.CATEGORY),
		Attributes: resources.CategoryAttributes{
			CategoryName: relateCategory.CategoryName,
			Unit:         relateCategory.Unit,
		},
	})

	result := resources.MealResponse{
		Data: resources.Meal{
			Key: resources.NewKeyInt64(resultMeal.Id, resources.MEAL),
			Attributes: resources.MealAttributes{
				MealName: resultMeal.MealName,
				Price:    resultMeal.Price,
				Amount:   resultMeal.Amount,
			},
			Relationships: resources.MealRelationships{
				Category: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultMeal.CategoryId, 10),
						Type: resources.CATEGORY,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
