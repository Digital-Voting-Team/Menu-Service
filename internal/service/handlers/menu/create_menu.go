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

func CreateMenu(endpointsConf *config.EndpointsConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := requests.NewCreateMenuRequest(r)
		if err != nil {
			helpers.Log(r).WithError(err).Info("wrong request")
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		Menu := data.Menu{
			CafeId: cast.ToInt64(request.Data.Relationships.Cafe.Data.ID),
		}

		var resultMenu data.Menu
		relateCafe, err := helpers.ValidateCafe(r.Header.Get("Authorization"), endpointsConf.Endpoints["cafe-service"], Menu.CafeId)
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to get cafe")
			ape.RenderErr(w, problems.NotFound())
			return
		}

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
