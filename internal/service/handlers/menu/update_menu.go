package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/menu"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"github.com/spf13/cast"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateMenuRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	menu, err := helpers.MenusQ(r).FilterById(request.MenuId).Get()
	if menu == nil {
		ape.Render(w, problems.NotFound())
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

	newMenu := data.Menu{
		CafeId: cast.ToInt64(resultCafe.Data.ID),
	}

	var resultMenu data.Menu
	resultMenu, err = helpers.MenusQ(r).FilterById(menu.Id).Update(newMenu)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to update menu")
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
