package mysql

import (
	"SolutionDev/internal/domain"
	"context"
)

type Video interface{
	GetStep(ctx context.Context, req *domain.Step) (*domain.Step, error)
	GetVideo(ctx context.Context, req *domain.Video) (*domain.Video, error)
}

type video struct{}

func (v *video) GetStep(ctx context.Context, req *domain.Step) (*domain.Step, error) {
	var step *domain.Step
	err := DBFromCtx(ctx).
		Where(&req).
		First(&step).Error
	if err != nil {
		return nil, err
	}
	return step, nil
}

func (v *video) GetVideo(ctx context.Context, req *domain.Video) (*domain.Video, error) {
	var video *domain.Video
	err := DBFromCtx(ctx).
		Where(&req).
		First(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}