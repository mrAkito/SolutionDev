package request

import "SolutionDev/internal/domain"

type Step struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type PostSteps struct {
	PreviousStep Step `json:"previous_step" validate:"required"`
	CurrentStep  Step `json:"current_step" validate:"required"`
	NextStep     Step `json:"next_step" validate:"required"`
}

type Video struct {
	CurrentStepId  int `json:"current_step_id" validate:"required"`
	PreviousStepId int `json:"previous_step_id" validate:"required"`
	NextStepId     int `json:"next_step_id" validate:"required"`
}

func ToStep(param Step) *domain.Step {
	return &domain.Step{
		Lat: param.Latitude,
		Lng: param.Longitude,
	}
}

func (param *PostSteps) ToSteps() []*domain.Step {
	var Steps []*domain.Step
	PreviousStep := param.PreviousStep
	CurrentStep := param.CurrentStep
	NextStep := param.NextStep
	Steps = append(Steps, ToStep(PreviousStep))
	Steps = append(Steps, ToStep(CurrentStep))
	Steps = append(Steps, ToStep(NextStep))
	return Steps
}

func ToVideo(stepIds []int) *domain.Video {
	return &domain.Video{
		PreviousStepId: stepIds[0],
		CurrentStepId:  stepIds[1],
		NextStepId:     stepIds[2],
	}
}
