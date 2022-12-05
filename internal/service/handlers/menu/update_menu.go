package handlers

import (
	"github.com/Digital-Voting-Team/menu-service/internal/config"
	"github.com/Digital-Voting-Team/menu-service/internal/data"
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	requests "github.com/Digital-Voting-Team/menu-service/internal/service/requests/menu"
	"github.com/Digital-Voting-Team/menu-service/resources"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func UpdateMenu(endpointsConf *config.EndpointsConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		newMenu := data.Menu{
			CafeId: cast.ToInt64(request.Data.Relationships.Cafe.Data.ID),
		}

		relateCafe, err := helpers.ValidateCafe(r.Header.Get("Authorization"), endpointsConf.Endpoints["cafe-service"], newMenu.CafeId)
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to get cafe")
			ape.RenderErr(w, problems.NotFound())
			return
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
							ID:   relateCafe.Data.ID,
							Type: resources.CAFE_REF,
						},
					},
				},
			},
		}
		ape.Render(w, result)
	}
}
