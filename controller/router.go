package controller

import (
	"context"
	"net/http"
	"strconv"

	"meh/core/user"

	"github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
)

const authHeader = "Authorization"

func NewRouter() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.Compress(5))
	r.Use(middleware.RequestID)
	r.Use(auth)
	return r, nil
}

func auth(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uidStr := r.Header.Get(authHeader)

		uid, err := strconv.ParseUint(uidStr, 10, 64)
		var ctx2 context.Context
		if err == nil {
			ctx2 = context.WithValue(ctx, user.IDContextKey, uid)
		} else {
			ctx2 = ctx
		}

		next.ServeHTTP(w, r.WithContext(ctx2))
	}

	return http.HandlerFunc(f)
}
