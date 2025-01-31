package models

import "time"

type Record struct {
	Date     time.Time
	Weight   float64
	Trained  bool
	Calories int
}
