package models

import "time"

type Board struct {
	BoardId      uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Id           string
	Name         string
	Day          string
	Totaltime    int
	Trytime      int
	Recoverytime int
	Frontcount   int
	Backcount    int
	AvgRPM       int
	AvgSpeed     float64
	Distance     float64
	Musclenum    float64
	Kcalorynum   float64
	Gender       string
	Area         string
	Birth        string
	Bike_info    string
	Career       string
	Club         string
	Email        string
}
