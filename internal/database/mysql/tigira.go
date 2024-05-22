package mysql

import (
	"SolutionDev/internal/domain"
	"context"
	"database/sql"
)

type Tigira interface {
	Create(ctx context.Context, req []*domain.GPSInfo) error
	GetGoalId(ctx context.Context, req *domain.Goal) (*domain.Goal, error)
	Update(ctx context.Context, req *domain.GPSInfo) error
	List(ctx context.Context, req *domain.GPSInfo) ([]*domain.GPSInfo, error)
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

func (t *tigira) Update(ctx context.Context, req *domain.GPSInfo) error {
	// 検索条件に一致するレコードのdeletedカラムの値を1に更新
	err := DBFromCtx(ctx).
		Model(&domain.GPSInfo{}).
		Where(req).
		Update("deleted", sql.NullInt64{Int64: 1, Valid: true}).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *tigira) List(ctx context.Context, req *domain.GPSInfo) ([]*domain.GPSInfo, error) {
	var res []*domain.GPSInfo
	err := DBFromCtx(ctx).
		Where(&req).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
