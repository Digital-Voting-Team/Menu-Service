package helpers

import (
	"Menu-Service/internal/data"
	"context"

	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	categoriesQCtxKey
	mealsQCtxKey
	receiptsQCtxKey
	menusQCtxKey
	mealMenusQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxCategoriesQ(entry data.CategoriesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, categoriesQCtxKey, entry)
	}
}

func CategoriesQ(r *http.Request) data.CategoriesQ {
	return r.Context().Value(categoriesQCtxKey).(data.CategoriesQ).New()
}

func CtxMealsQ(entry data.MealsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mealsQCtxKey, entry)
	}
}

func MealsQ(r *http.Request) data.MealsQ {
	return r.Context().Value(mealsQCtxKey).(data.MealsQ).New()
}

func CtxReceiptsQ(entry data.ReceiptsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, receiptsQCtxKey, entry)
	}
}

func ReceiptsQ(r *http.Request) data.ReceiptsQ {
	return r.Context().Value(receiptsQCtxKey).(data.ReceiptsQ).New()
}

func CtxMenusQ(entry data.MenusQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, menusQCtxKey, entry)
	}
}

func MenusQ(r *http.Request) data.MenusQ {
	return r.Context().Value(menusQCtxKey).(data.MenusQ).New()
}

func CtxMealMenusQ(entry data.MealMenusQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mealMenusQCtxKey, entry)
	}
}

func MealMenusQ(r *http.Request) data.MealMenusQ {
	return r.Context().Value(mealMenusQCtxKey).(data.MealMenusQ).New()
}
