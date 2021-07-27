package helpers

import "time"

type Currency	struct{
	Name	string
	Code	string
	Multiplier	float64	// base is PLN
	Date	time.Time
}