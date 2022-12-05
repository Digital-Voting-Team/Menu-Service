package handlers

import (
	"github.com/spf13/cast"
	"menu-service/internal/data"
	"menu-service/internal/service/helpers"
	requests "menu-service/internal/service/requests/menu"
	"menu-service/resources"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateMenuRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	cafeId := cast.ToInt64(request.Data.Relationships.Cafe.Data.ID)
	resultCafe, err := helpers.ValidateCafe(r.Header.Get("Authorization"), r.Context().Value("cafeEndpoint").(string), cafeId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get cafe from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if resultCafe == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	Menu := data.Menu{
		CafeId: cast.ToInt64(resultCafe.Data.ID),
	}

	var resultMenu data.Menu

	resultMenu, err = helpers.MenusQ(r).Insert(Menu)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create menu")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.MenuResponse{
		Data: resources.Menu{
			Key: resources.NewKeyInt64(resultMenu.Id, resources.MENU),
			Relationships: resources.MenuRelationships{
				Cafe: resources.Relation{
					Data: &resources.Key{
						ID:   strconv.FormatInt(resultMenu.CafeId, 10),
						Type: resources.CAFE_REF,
					},
				},
			},
		},
	}
	ape.Render(w, result)
}
