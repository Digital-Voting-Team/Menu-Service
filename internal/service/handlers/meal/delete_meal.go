package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/meal"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteMeal(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteMealRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Meal, err := helpers.MealsQ(r).FilterById(request.MealId).Get()
	if Meal == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.MealsQ(r).Delete(request.MealId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete meal")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
