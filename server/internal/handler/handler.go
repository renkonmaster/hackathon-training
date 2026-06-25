package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/renkonmaster/hackathon-training/server/internal/gen"
	"github.com/renkonmaster/hackathon-training/server/internal/service"
)

type Handler struct {
	service *service.Service
}

var _ gen.ServerInterface = (*Handler)(nil)

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetApi(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (h *Handler) GetApiLang(ctx echo.Context, params gen.GetApiLangParams) error {
	var lang string
	if params.Lang == nil {
		lang = "en"
	} else {
		lang = *params.Lang
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": h.service.LanguageMessage(lang),
	})
}

func (h *Handler) GetPing(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": h.service.PingMessage(),
	})
}
