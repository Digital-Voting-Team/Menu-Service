package handlers

import (
	"menu-service/internal/service/helpers"
	requests "menu-service/internal/service/requests/category"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteCategoryRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	category, err := helpers.CategoriesQ(r).FilterById(request.CategoryId).Get()
	if category == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.CategoriesQ(r).Delete(request.CategoryId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete category")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
