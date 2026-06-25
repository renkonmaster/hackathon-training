package service

import (
	"strings"

	"github.com/renkonmaster/hackathon-training/server/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) PingMessage() string {
	return "pong"
}

func (s *Service) LanguageMessage(lang string) string {
	switch strings.ToLower(lang) {
	case "ja":
		return "こんにちは"
	case "en":
		return "Hello"
	default:
		return "Hello"
	}
}
