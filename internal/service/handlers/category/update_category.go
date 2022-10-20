package handlers

import (
	"Menu-Service/internal/data"
	"Menu-Service/internal/service/helpers"
	requests "Menu-Service/internal/service/requests/category"
	"Menu-Service/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateCategoryRequest(r)
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

	newCategory := data.Category{
		CategoryName: request.Data.Attributes.CategoryName,
		Unit:         request.Data.Attributes.Unit,
	}

	var resultCategory data.Category
	resultCategory, err = helpers.CategoriesQ(r).FilterById(category.Id).Update(newCategory)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update category")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.CategoryResponse{
		Data: resources.Category{
			Key: resources.NewKeyInt64(resultCategory.Id, resources.CATEGORY),
			Attributes: resources.CategoryAttributes{
				CategoryName: resultCategory.CategoryName,
				Unit:         resultCategory.Unit,
			},
		},
	}
	ape.Render(w, result)
}
