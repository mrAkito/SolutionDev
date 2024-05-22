package request

import "SolutionDev/internal/domain"

type Step struct {
	latitude	float64 `json:"latitude" validate:"required"`
	longitude	float64 `json:"longitude" validate:"required"`
}

type PostSteps struct {
	PreviousStep Step `json:"" validate:"required"`
	CurrentStep  Step `json:"" validate:"required"`
	NextStep     Step `json:"" validate:"required"`
}

type Video struct {
	CurrentStepId 	int `json:"current_step_id" validate:"required"`
	PreviousStepId 	int `json:"previous_step_id" validate:"required"`
	NextStepId 			int `json:"next_step_id" validate:"required"`
}

func ToStep(param Step) *domain.Step {
	return &domain.Step{
		Lat: param.latitude,
		Lng: param.longitude,
	}
}

func (param *PostSteps) ToSteps() []*domain.Step {
	var Steps []*domain.Step
	PreviousStep :=  param.PreviousStep
	CurrentStep := param.CurrentStep
	NextStep := param.NextStep
	Steps = append(Steps, ToStep(PreviousStep))
	Steps = append(Steps, ToStep(CurrentStep))
	Steps = append(Steps, ToStep(NextStep))
	return Steps
}

func ToVideo(stepIds []int) *domain.Video {
	return &domain.Video{
		CurrentStepId: stepIds[0],
		PreviousStepId: stepIds[1],
		NextStepId: stepIds[2],
	}
}