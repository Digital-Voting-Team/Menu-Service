package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/menu"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetMenuList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetMenuListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	menusQ := helpers.MenusQ(r)
	applyFilters(menusQ, request)
	menus, err := menusQ.Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get menus")
		ape.Render(w, problems.InternalError())
		return
	}

	response := resources.MenuListResponse{
		Data:  newMenusList(menus),
		Links: helpers.GetOffsetLinks(r, request.OffsetPageParams),
	}
	ape.Render(w, response)
}

func applyFilters(q data.MenusQ, request requests.GetMenuListRequest) {
	q.Page(request.OffsetPageParams)

	if len(request.FilterCafeId) > 0 {
		q.FilterByCafeId(request.FilterCafeId...)
	}
}

func newMenusList(menus []data.Menu) []resources.Menu {
	result := make([]resources.Menu, len(menus))
	for i, menu := range menus {
		result[i] = resources.Menu{
			Key: resources.NewKeyInt64(menu.Id, resources.MENU),
			Relationships: resources.MenuRelationships{
				Cafe: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(menu.CafeId, 10),
						Type: resources.CAFE_REF,
					},
				},
			},
		}
	}
	return result
}
