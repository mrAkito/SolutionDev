package main

import (
	"fmt"
	"gorm.io/gorm"
	"SolutionDev/internal/database/mysql"
	"SolutionDev/internal/domain"
)

func seeds(gdb *gorm.DB) error {
	goal := domain.Goal{ GoalName: "ドンキ・ホーテ" }

	if err := gdb.Create(&goal).Error; err != nil {
		fmt.Println("%+v", err)
		return err
	}

	steps := []domain.Step{
		{ Lat:  35.6917081, Lng: 139.6966017 },
		{ Lat:  35.69295690000001, Lng: 139.6968302 },
		{ Lat:  35.6930854, Lng: 139.6967737 },
		{ Lat:  35.69372329999999, Lng: 139.6947911 },
		{ Lat:  35.6943641, Lng: 139.69524 },
		{ Lat:  35.6945804, Lng: 139.6947985 },
		{ Lat:  35.6950534, Lng: 139.6950874 },
		{ Lat:  35.6951818, Lng: 139.6947902 },
	}

	if err := gdb.Create(&steps).Error; err != nil {
		fmt.Println("%+v", err)
		return err
	}

	videos := []domain.Video{
		{ CurrentStepId: 2, PreviousStepId: 1, NextStepId: 3, Url: "https://www.youtube.com/watch?v=2" },
		{ CurrentStepId: 3, PreviousStepId: 2, NextStepId: 4, Url: "https://www.youtube.com/watch?v=3" },
		{ CurrentStepId: 5, PreviousStepId: 4, NextStepId: 6, Url: "https://www.youtube.com/watch?v=5" },
		{ CurrentStepId: 6, PreviousStepId: 5, NextStepId: 7, Url: "https://www.youtube.com/watch?v=6" },
		{ CurrentStepId: 7, PreviousStepId: 6, NextStepId: 8, Url: "https://www.youtube.com/watch?v=7" },
		{ CurrentStepId: 8, PreviousStepId: 7, NextStepId: 9, Url: "https://www.youtube.com/watch?v=8" },
		{ CurrentStepId: 9, PreviousStepId: 8, NextStepId: 10, Url: "https://www.youtube.com/watch?v=9" },
	}

	if err := gdb.Create(&videos).Error; err != nil {
		fmt.Println("%+v", err)
		return err
	}
	return nil
}

func main() {
	gdb, err := mysql.Connect()
	if err != nil {
		return
	}

	defer mysql.Close(gdb)

	if err := seeds(gdb); err != nil {
		fmt.Println("%+v", err)
		return
	}
}