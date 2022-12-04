package handlers

import (
	"github.com/spf13/cast"
	"menu-service/internal/data"
	"menu-service/internal/service/helpers"
	requests "menu-service/internal/service/requests/receipt"
	"menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateReceiptRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Receipt := data.Receipt{
		Quantity:     request.Data.Attributes.Quantity,
		MealId:       cast.ToInt64(request.Data.Relationships.Meal.Data.ID),
		IngredientId: cast.ToInt64(request.Data.Relationships.Ingredient.Data.ID),
	}

	var resultReceipt data.Receipt
	relateMeal, err := helpers.MealsQ(r).FilterById(Receipt.MealId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get meal")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	resultReceipt, err = helpers.ReceiptsQ(r).Insert(Receipt)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create receipt")
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

	result := resources.ReceiptResponse{
		Data: resources.Receipt{
			Key: resources.NewKeyInt64(resultReceipt.Id, resources.RECEIPT),
			Attributes: resources.ReceiptAttributes{
				Quantity: resultReceipt.Quantity,
			},
			Relationships: resources.ReceiptRelationships{
				Meal: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultReceipt.MealId, 10),
						Type: resources.MEAL,
					},
				},
				Ingredient: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultReceipt.IngredientId, 10),
						Type: resources.INGREDIENT_REF,
					},
				},
			},
		},
		Included: includes,
	}
	ape.Render(w, result)
}
