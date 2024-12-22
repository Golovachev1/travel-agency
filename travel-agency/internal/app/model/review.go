package model

import "time"

type Review struct {
	ID               int
    Score int
	Review_text string
	Publish_date time.Time
    User User
	Tour           Tour
}