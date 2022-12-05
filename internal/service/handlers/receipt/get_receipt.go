package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/receipt"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetReceipt(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetReceiptRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resultReceipt, err := helpers.ReceiptsQ(r).FilterById(request.ReceiptId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get receipt from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultReceipt == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	relateMeal, err := helpers.MealsQ(r).FilterById(resultReceipt.MealId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get meal")
		ape.RenderErr(w, problems.NotFound())
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
