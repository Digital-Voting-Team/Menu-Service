package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/menu"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteMenu(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteMenuRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Menu, err := helpers.MenusQ(r).FilterById(request.MenuId).Get()
	if Menu == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.MenusQ(r).Delete(request.MenuId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete menu")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
