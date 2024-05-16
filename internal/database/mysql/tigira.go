package mysql

import (
	"SolutionDev/internal/domain"
	"context"
)

type Tigira interface {
	Create(ctx context.Context, req []*domain.GPSInfo) error
	GetGoalId(ctx context.Context, req *domain.Goal) (*domain.Goal, error)
}

type tigira struct{}

func (t *tigira) Create(ctx context.Context, req []*domain.GPSInfo) error {
	err := DBFromCtx(ctx).
		CreateInBatches(req, len(req)).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *tigira) GetGoalId(ctx context.Context, req *domain.Goal) (*domain.Goal, error) {
	var goal *domain.Goal
	err := DBFromCtx(ctx).
		Where(&req).
		First(&goal).Error
	if err != nil {
		return nil, err
	}
	return goal, nil
}
