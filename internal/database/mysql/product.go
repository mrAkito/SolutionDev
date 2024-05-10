package mysql

import (
	"SolutionDev/internal/domain"
	"context"
)

type Product interface {
	List(ctx context.Context) ([]*domain.Product, error)
	Get(ctx context.Context, req *domain.Product) (*domain.Product, error)
}

type product struct {
}

func (p *product) List(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	err := DBFromCtx(ctx).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *product) Get(ctx context.Context, req *domain.Product) (*domain.Product, error) {
	var getProduct *domain.Product
	err := DBFromCtx(ctx).
		Where(&req).
		First(&getProduct).Error
	if err != nil {
		return nil, err
	}
	return getProduct, nil
}
