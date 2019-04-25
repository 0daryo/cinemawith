package main

import (
	"net/http"

	"github.com/0daryo/cinemawith/src/config"
	"github.com/0daryo/cinemawith/src/handler"
	"github.com/0daryo/cinemawith/src/lib/log"
	"github.com/0daryo/cinemawith/src/middleware"
	"github.com/go-chi/chi"
)

// Routing ... ルーティング設定
func Routing(r chi.Router, d Dependency) {

	// request log
	r.Use(log.Middleware)

	// アクセスコントロール
	r.Use(middleware.AccessControl)

	// 認証なし(Stagingのみ)
	if config.IsEnvDeveloping() {
		r.Route("/noauth/v1", func(r chi.Router) {
			r.Use(d.DummyFirebaseAuth.Handle)
			r.Use(d.DummyHTTPHeader.Handle)
			subRouting(r, d)
		})
	}

	// 認証あり
	r.Route("/v1", func(r chi.Router) {
		r.Use(d.FirebaseAuth.Handle)
		r.Use(d.HTTPHeader.Handle)
		subRouting(r, d)
	})

	// Ping
	r.Get("/ping", handler.Ping)
	r.Get("/", handler.Ping)

	http.Handle("/", r)
}

func subRouting(r chi.Router, d Dependency) {
	// API
	r.Route("/users", func(r chi.Router) {
		//r.Post("/", d.UserHandler.Create)
		r.Get("/{userID}", d.UserHandler.Get)
	})

}
