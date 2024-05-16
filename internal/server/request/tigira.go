package request

import "SolutionDev/internal/domain"

type GetGoal struct {
	GoalName string `param:"goal_name" validate:"required"`
}

type GPSInfo struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type PostGPSInfo struct {
	GoalName string    `param:"goal_name" validate:"required"`
	GPSInfos []GPSInfo `json:"" validate:"required,dive"`
}

func (param *GetGoal) ToGoal() *domain.Goal {
	return &domain.Goal{
		GoalName: param.GoalName,
	}
}

func ToGPSInfo(GPS GPSInfo, id int) *domain.GPSInfo {
	return &domain.GPSInfo{
		Id:  id,
		Lat: GPS.Latitude,
		Lng: GPS.Longitude,
	}
}

func (param *PostGPSInfo) ToGPSInfos(id int) []*domain.GPSInfo {
	var GPSinfos []*domain.GPSInfo
	for _, p := range param.GPSInfos {
		GPSinfos = append(GPSinfos, ToGPSInfo(p, id))
	}
	return GPSinfos
}

func (param *PostGPSInfo) ToGPSInfoGoalId() *domain.Goal {
	return &domain.Goal{
		GoalName: param.GoalName,
	}
}
