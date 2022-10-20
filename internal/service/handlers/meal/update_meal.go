package handlers

import (
	"Menu-Service/internal/data"
	"Menu-Service/internal/service/helpers"
	requests "Menu-Service/internal/service/requests/meal"
	"Menu-Service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateMeal(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateMealRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	meal, err := helpers.MealsQ(r).FilterById(request.MealId).Get()
	if meal == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	newMeal := data.Meal{
		MealName:   request.Data.Attributes.MealName,
		Price:      request.Data.Attributes.Price,
		Amount:     request.Data.Attributes.Amount,
		CategoryId: cast.ToInt64(request.Data.Relationships.Category.Data.ID),
	}

	relateCategory, err := helpers.CategoriesQ(r).FilterById(newMeal.CategoryId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get new category")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	var resultMeal data.Meal
	resultMeal, err = helpers.MealsQ(r).FilterById(meal.Id).Update(newMeal)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update meal")
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
