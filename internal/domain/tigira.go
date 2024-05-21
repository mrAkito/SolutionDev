package domain

import "database/sql"

type GPSInfo struct {
	Id  int            `gorm:"column:id"`
	Lat float64        `gorm:"column:longitude"`
	Lng float64        `gorm:"column:latitude"`
	Opt sql.NullString `gorm:"column:option"`
	Dlt sql.NullInt64  `gorm:"column:deleted"`
}

type Goal struct {
	GoalId   int    `gorm:"column:goal_id"`
	GoalName string `gorm:"column:goal_name"`
}

func (GPSInfo) TableName() string {
	return "GPSInfo"
}

func (Goal) TableName() string {
	return "Goal"
}
