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
				helpers.Log(r).Info("insufficient user permissions (Manager)")
				ape.RenderErr(w, problems.Forbidden())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func CheckAdminPosition() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
			if accessLevel < staffRes.Admin {
				helpers.Log(r).Info("insufficient user permissions (Administrator)")
				ape.RenderErr(w, problems.Forbidden())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func CheckAccountantPosition() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
			if accessLevel < staffRes.Accountant {
				helpers.Log(r).Info("insufficient user permissions (Accountant)")
				ape.RenderErr(w, problems.Forbidden())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func CheckWorkerPosition() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessLevel := r.Context().Value("accessLevel").(staffRes.AccessLevel)
			if accessLevel < staffRes.Worker {
				helpers.Log(r).Info("insufficient user permissions (Worker)")
				ape.RenderErr(w, problems.Forbidden())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
