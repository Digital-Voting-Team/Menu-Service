package handlers

import (
	"menu-service/internal/service/helpers"
	requests "menu-service/internal/service/requests/meal"
	"menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetMeal(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetMealRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultMeal, err := helpers.MealsQ(r).FilterById(request.MealId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get meal from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultMeal == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	relateCategory, err := helpers.CategoriesQ(r).FilterById(resultMeal.CategoryId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get category")
		ape.RenderErr(w, problems.NotFound())
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
