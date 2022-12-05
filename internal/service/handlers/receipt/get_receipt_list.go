package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/receipt"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetReceiptList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetReceiptListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	receiptsQ := helpers.ReceiptsQ(r)
	applyFilters(receiptsQ, request)
	receipts, err := receiptsQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get receipts")
		ape.Render(w, problems.InternalError())
		return
	}
	meals, err := helpers.MealsQ(r).FilterById(getMealIds(receipts)...).Select()

	response := resources.ReceiptListResponse{
		Data:     newReceiptsList(receipts),
		Links:    helpers.GetOffsetLinks(r, request.OffsetPageParams),
		Included: newReceiptIncluded(meals),
	}
	ape.Render(w, response)
}

func applyFilters(q data.ReceiptsQ, request requests.GetReceiptListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterMealId) > 0 {
		q.FilterByMealId(request.FilterMealId...)
	}

	if len(request.FilterIngredientId) > 0 {
		q.FilterByIngredientId(request.FilterIngredientId...)
	}

	if len(request.FilterQuantityFrom) > 0 {
		q.FilterByQuantityFrom(request.FilterQuantityFrom...)
	}

	if len(request.FilterQuantityTo) > 0 {
		q.FilterByQuantityTo(request.FilterQuantityTo...)
	}
}

func newReceiptsList(receipts []data.Receipt) []resources.Receipt {
	result := make([]resources.Receipt, len(receipts))
	for i, receipt := range receipts {
		result[i] = resources.Receipt{
			Key: resources.NewKeyInt64(receipt.Id, resources.RECEIPT),
			Attributes: resources.ReceiptAttributes{
				Quantity: receipt.Quantity,
			},
			Relationships: resources.ReceiptRelationships{
				Meal: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(receipt.MealId, 10),
						Type: resources.MEAL,
					},
				},
				Ingredient: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(receipt.IngredientId, 10),
						Type: resources.INGREDIENT_REF,
					},
				},
			},
		}
	}
	return result
}

func getMealIds(receipts []data.Receipt) []int64 {
	mealIDs := make([]int64, len(receipts))
	for i := 0; i < len(receipts); i++ {
		mealIDs[i] = receipts[i].MealId
	}
	return mealIDs
}

func newReceiptIncluded(meals []data.Meal) resources.Included {
	result := resources.Included{}
	for _, item := range meals {
		resource := newMealModel(item)
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
