package main

import (
	"github.com/0daryo/cinemawith/src/handler/api"
	"github.com/0daryo/cinemawith/src/infrastructure/repository"
	"github.com/0daryo/cinemawith/src/lib/firebaseauth"
	"github.com/0daryo/cinemawith/src/lib/httpheader"
	"github.com/0daryo/cinemawith/src/lib/mysql"
	"github.com/0daryo/cinemawith/src/service"
)

// Dependency ... dependency
type Dependency struct {
	DummyFirebaseAuth *firebaseauth.Middleware
	FirebaseAuth      *firebaseauth.Middleware
	DummyHTTPHeader   *httpheader.Middleware
	HTTPHeader        *httpheader.Middleware
	UserHandler       *api.UserHandler
}

// Inject ... indect dependency
func (d *Dependency) Inject() {
	// Config
	dbCfg := mysql.GetSQLConfig()

	// Lib
	dbConn := mysql.NewSQLClient(dbCfg)

	// Repository
	uRepo := repository.NewUser(dbConn)

	// Service
	dfaSvc := firebaseauth.NewDummyService()
	faSvc := firebaseauth.NewService()
	dhhSvc := httpheader.NewDummyService()
	hhSvc := httpheader.NewService()
	uSvc := service.NewUser(uRepo)

	// Middleware
	d.DummyFirebaseAuth = firebaseauth.NewMiddleware(dfaSvc)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)
	d.DummyHTTPHeader = httpheader.NewMiddleware(dhhSvc)
	d.HTTPHeader = httpheader.NewMiddleware(hhSvc)

	// Handler
	d.UserHandler = api.NewUserHandler(uSvc)
}
