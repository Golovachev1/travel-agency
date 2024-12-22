package model

import "time"

type Tour_Base struct {
	ID               int
	City_Of_Flight string
    Arrival_Country string
    Duration int
    Date_Of_Tour time.Time
    Tour_Cost float64
	Tour           Tour
}