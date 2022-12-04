package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"menu-service/internal/data/pg"
	category "menu-service/internal/service/handlers/category"
	meal "menu-service/internal/service/handlers/meal"
	mealMenu "menu-service/internal/service/handlers/meal_menu"
	menu "menu-service/internal/service/handlers/menu"
	receipt "menu-service/internal/service/handlers/receipt"
	"menu-service/internal/service/middleware"

	"menu-service/internal/service/helpers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	log := s.log.WithFields(map[string]interface{}{
		"service": "menu-service",
	})

	r.Use(
		ape.RecoverMiddleware(log),
		ape.LoganMiddleware(log),
		ape.CtxMiddleware(
			helpers.CtxLog(log),
			helpers.CtxCategoriesQ(pg.NewCategoriesQ(s.db)),
			helpers.CtxMealsQ(pg.NewMealsQ(s.db)),
			helpers.CtxReceiptsQ(pg.NewReceiptsQ(s.db)),
			helpers.CtxMenusQ(pg.NewMenusQ(s.db)),
			helpers.CtxMealMenusQ(pg.NewMealMenusQ(s.db)),
		),
		middleware.BasicAuth(s.endpoints),
	)
	r.Route("/integrations/menu-service", func(r chi.Router) {
		r.Use(middleware.CheckManagerPosition())
		r.Route("/categories", func(r chi.Router) {
			r.Post("/", category.CreateCategory)
			r.Get("/", category.GetCategoryList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", category.GetCategory)
				r.Put("/", category.UpdateCategory)
				r.Delete("/", category.DeleteCategory)
			})
		})
		r.Route("/meals", func(r chi.Router) {
			r.Post("/", meal.CreateMeal)
			r.Get("/", meal.GetMealList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", meal.GetMeal)
				r.Put("/", meal.UpdateMeal)
				r.Delete("/", meal.DeleteMeal)
			})
		})
		r.Route("/receipts", func(r chi.Router) {
			r.Post("/", receipt.CreateReceipt)
			r.Get("/", receipt.GetReceiptList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", receipt.GetReceipt)
				r.Put("/", receipt.UpdateReceipt)
				r.Delete("/", receipt.DeleteReceipt)
			})
		})
		r.Route("/menus", func(r chi.Router) {
			r.Post("/", menu.CreateMenu)
			r.Get("/", menu.GetMenuList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", menu.GetMenu)
				r.Put("/", menu.UpdateMenu)
				r.Delete("/", menu.DeleteMenu)
			})
		})
		r.Route("/meal_menus", func(r chi.Router) {
			r.Post("/", mealMenu.CreateMealMenu)
			r.Get("/", mealMenu.GetMealMenuList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", mealMenu.GetMealMenu)
				r.Put("/", mealMenu.UpdateMealMenu)
				r.Delete("/", mealMenu.DeleteMealMenu)
			})
		})
	})

	return r
}
