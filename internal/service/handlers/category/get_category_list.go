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

func GetCategoryList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetCategoryListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	categoriesQ := helpers.CategoriesQ(r)
	applyFilters(categoriesQ, request)
	category, err := categoriesQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get category")
		ape.Render(w, problems.InternalError())
		return
	}

	response := resources.CategoryListResponse{
		Data:  newCategoriesList(category),
		Links: helpers.GetOffsetLinks(r, request.OffsetPageParams),
	}
	ape.Render(w, response)
}

func applyFilters(q data.CategoriesQ, request requests.GetCategoryListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterCategoryName) > 0 {
		q.FilterByNames(request.FilterCategoryName...)
	}

	if len(request.FilterUnit) > 0 {
		q.FilterByUnits(request.FilterUnit...)
	}
}

func newCategoriesList(categories []data.Category) []resources.Category {
	result := make([]resources.Category, len(categories))
	for i, category := range categories {
		result[i] = resources.Category{
			Key: resources.NewKeyInt64(category.Id, resources.CATEGORY),
			Attributes: resources.CategoryAttributes{
				CategoryName: category.CategoryName,
				Unit:         category.Unit,
			},
		}
	}
	return result
}
