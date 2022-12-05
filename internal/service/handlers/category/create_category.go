package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/category"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateCategoryRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var resultCategory data.Category

	category := data.Category{
		CategoryName: request.Data.Attributes.CategoryName,
		Unit:         request.Data.Attributes.Unit,
	}

	resultCategory, err = helpers.CategoriesQ(r).Insert(category)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create category")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.CategoryResponse{
		Data: resources.Category{
			Key: resources.NewKeyInt64(resultCategory.Id, resources.CATEGORY),
			Attributes: resources.CategoryAttributes{
				CategoryName: request.Data.Attributes.CategoryName,
				Unit:         request.Data.Attributes.Unit,
			},
		},
	}
	ape.Render(w, result)
}
