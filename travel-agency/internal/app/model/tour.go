package model

import "time"

type Tour struct {
	ID               int
	Start_Date       time.Time
	End_Date         time.Time
	Amount_Of_People string
	User            User
}