package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	mw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		mw.Logger,
	)
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
