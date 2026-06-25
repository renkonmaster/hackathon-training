package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/renkonmaster/hackathon-training/server/internal/gen"
	"github.com/renkonmaster/hackathon-training/server/internal/handler"
	"github.com/renkonmaster/hackathon-training/server/internal/repository"
	"github.com/renkonmaster/hackathon-training/server/internal/service"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderContentType},
		AllowOrigins: []string{"*"},
	}))

	repo := repository.New()
	s := service.New(repo)
	h := handler.New(s)
	gen.RegisterHandlers(e, h)

	addr := getEnv("API_ADDR", ":"+getEnv("PORT", "8080"))
	if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		e.Logger.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
