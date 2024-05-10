package request

import "SolutionDev/internal/domain"

type GetProduct struct {
	Id string `param:"id" validate:"required,uuid"`
}

func (param *GetProduct) ToProduct() *domain.Product {
	return &domain.Product{
		Id: param.Id,
	}
}
