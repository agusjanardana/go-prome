package main

import (
	"log"
	"net/http"

	notecontroller "github.com/go-prome/app/controllers/note_controller"
	noterepository "github.com/go-prome/app/drivers/repository/note_repository"
	middlewares "github.com/go-prome/app/middleware"
	noteservices "github.com/go-prome/app/services/note_services"
	configs "github.com/go-prome/config"
	"github.com/go-prome/config/db"
	v1 "github.com/go-prome/router/v1"
	"github.com/go-prome/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10, // 1 KB
		DisableStackAll: true,
	}))
	e.HTTPErrorHandler = utils.ErrorHandler

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderCookie, echo.HeaderAccessControlAllowCredentials, echo.HeaderAccessControlAllowOrigin},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middlewares.MetricsMiddleware)

	cfg := configs.New()

	databaseClient := db.New(cfg)
	defer databaseClient.Close()

	// repository
	noteRepository := noterepository.NewNoteRepositoryImpl(databaseClient)

	// services
	noteServices := noteservices.NewNoteServices(noteRepository)

	// controller
	noteController := notecontroller.NewNoteControllerImpl(noteServices)

	routeV1init := v1.ControllerList{
		NoteController: noteController,
	}

	routeV1init.Registration(e)

	// metrics
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	if err := e.Start(":" + cfg.Get("PORT")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
