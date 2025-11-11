package service

import (
	"context"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/repository"
)

type HistoryService struct {
	repo *repository.HistoryRepository
}

func NewHistoryService(repo *repository.HistoryRepository) *HistoryService {
	return &HistoryService{repo: repo}
}

func (s *HistoryService) GetHistoryByGlassesID(ctx context.Context, id uint) ([]model.StatusHistoryResponse, error) {
	return s.repo.FindHistoryByGlassesID(ctx, id)
}

func (s *HistoryService) CreateHistory(ctx context.Context, history *model.StatusHistory) error {
	return s.repo.CreateHistory(ctx, history)
}
