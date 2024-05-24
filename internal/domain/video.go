package domain

type Step struct {
	Id 			int     `gorm:"column:id"`
	Lat 		float64 `gorm:"column:latitude"`
	Lng 		float64 `gorm:"column:longitude"`
}

type Video struct {
	Id 							int 	 `gorm:"column:id"`
	CurrentStepId 	int 	 `gorm:"column:current_step_id"`
	PreviousStepId 	int 	 `gorm:"column:previous_step_id"`
	NextStepId 			int 	 `gorm:"column:next_step_id"`
	Url 						string `gorm:"column:url"`
}

func (Step) TableName() string {
	return "Steps"
}

func (Video) TableName() string {
	return "Videos"
}