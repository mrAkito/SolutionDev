package response

import (
	"SolutionDev/internal/domain"
)

type Goal struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ToGoals(products []*domain.GPSInfo) []*Goal {
	var res []*Goal
	for _, p := range products {
		res = append(res, ToGoal(p))
	}

	return res
}

func ToGoal(goal *domain.GPSInfo) *Goal {
	return &Goal{
		Latitude:  goal.Lat,
		Longitude: goal.Lng,
	}
}
