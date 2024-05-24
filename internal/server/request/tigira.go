package request

import (
	"SolutionDev/internal/conversion/null"
	"SolutionDev/internal/domain"
)

type GPSInfo struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type PostGPSInfo struct {
	GoalName string    `param:"goal_name" validate:"required"`
	Option   *string   `json:"option" validate:"omitempty"`
	GPSInfos []GPSInfo `json:"" validate:"required,dive"`
}

type GetGPSInfo struct {
	GoalName string  `query:"goal_name" validate:"required"`
	Option   *string `query:"option" validate:"omitempty"`
}

func ToGPSInfo(gps GPSInfo, id int, opt *string) *domain.GPSInfo {
	return &domain.GPSInfo{
		Id:  id,
		Lat: gps.Latitude,
		Lng: gps.Longitude,
		Opt: null.NullFromPtrStringOpt(opt),
		Dlt: null.NullFromPtrInt(nil),
	}
}

func (param *PostGPSInfo) ToSearchGPSInfo(id int) *domain.GPSInfo {
	return &domain.GPSInfo{
		Id:  id,
		Opt: null.NullFromPtrStringOpt(param.Option),
		Dlt: null.NullFromPtrInt(nil),
	}
}

func (param *PostGPSInfo) ToGPSInfos(id int) []*domain.GPSInfo {
	var GPSinfos []*domain.GPSInfo
	for _, p := range param.GPSInfos {
		GPSinfos = append(GPSinfos, ToGPSInfo(p, id, param.Option))
	}
	return GPSinfos
}

func (param *PostGPSInfo) ToGPSInfoGoalId() *domain.Goal {
	return &domain.Goal{
		GoalName: param.GoalName,
	}
}

func (param *GetGPSInfo) ToGPSInfo() *domain.Goal {
	return &domain.Goal{
		GoalName: param.GoalName,
	}
}

func (param *GetGPSInfo) ToGoal(id int, opt *string) *domain.GPSInfo {
	return &domain.GPSInfo{
		Id:  id,
		Opt: null.NullFromPtrStringOpt(opt),
		Dlt: null.NullFromPtrInt(nil),
	}
}
