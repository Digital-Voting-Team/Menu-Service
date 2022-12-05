package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/category"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetCategoryRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	category, err := helpers.CategoriesQ(r).FilterById(request.CategoryId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get category from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if category == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	result := resources.CategoryResponse{
		Data: resources.Category{
			Key: resources.NewKeyInt64(category.Id, resources.CATEGORY),
			Attributes: resources.CategoryAttributes{
				CategoryName: category.CategoryName,
				Unit:         category.Unit,
			},
		},
	}

	ape.Render(w, result)
}
