package model

import "time"

type Reservation struct {
	ID               int
	Reservation_Date       time.Time
	Payment float64
	Tour            Tour
}