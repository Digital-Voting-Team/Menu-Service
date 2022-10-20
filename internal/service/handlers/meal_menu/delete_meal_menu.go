package handlers

import (
	"Menu-Service/internal/service/helpers"
	requests "Menu-Service/internal/service/requests/meal_menu"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteMealMenu(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteMealMenuRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	MealMenu, err := helpers.MealMenusQ(r).FilterById(request.MealMenuId).Get()
	if MealMenu == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.MealMenusQ(r).Delete(request.MealMenuId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete mealMenu")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
