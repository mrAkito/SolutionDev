package response

import (
	"SolutionDev/internal/conversion/ptr"
	"SolutionDev/internal/domain"
	"time"
)

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Version    int       `json:"version"`
	Prefecture *string   `json:"prefecture,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func ToProducts(products []*domain.Product) []*Product {
	var res []*Product
	for _, p := range products {
		res = append(res, ToProduct(p))
	}

	return res
}

func ToProduct(product *domain.Product) *Product {
	return &Product{
		ID:         product.Id,
		Name:       product.Name,
		Version:    product.Version,
		Prefecture: ptr.PtrFromString(product.Prefecture),
		CreatedAt:  product.CreatedAt,
	}
}
