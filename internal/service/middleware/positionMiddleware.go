package middleware

import (
	"github.com/Digital-Voting-Team/menu-service/internal/service/helpers"
	staffRes "github.com/Digital-Voting-Team/staff-service/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"net/http"
)

func CheckManagerPosition() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
			if accessLevel < staffRes.Manager {
				helpers.Log(r).Info("insufficient user permissions")
				ape.RenderErr(w, problems.Forbidden())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
