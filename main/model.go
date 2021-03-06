package main

import "time"

var employees = map[string]Employee{
	"962145": Employee{
		ID:        962145,
		FirstName: "first name 1",
		LastName:  "last name 1",
		Position:  "position 1",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		TotalPTO:  30.,
	},
	"962146": Employee{
		ID:        962146,
		FirstName: "first name 2",
		LastName:  "last name 2",
		Position:  "position 2",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		TotalPTO:  30.,
	},
	"962147": Employee{
		ID:        962147,
		FirstName: "first name 3",
		LastName:  "last name 3",
		Position:  "position 3",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		TotalPTO:  30.,
	},
	"962148": Employee{
		ID:        962148,
		FirstName: "first name 4",
		LastName:  "last name 4",
		Position:  "position 4",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		TotalPTO:  30.,
	},
}

var TimesOff = map[string][]TimeOff{
	"962145": []TimeOff{
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2016, 12, 8, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}

type Employee struct {
	ID        int
	TotalPTO  float32 `form:"pto" json:"pto"`
	FirstName string  `form:"firstName" json:"firstName"`
	LastName  string  `form:"lastName" json:"lastName"`
	Position  string  `form:"position" json:"position"`
	Status    string
	StartDate time.Time
	TimesOff  []TimeOff
}

type TimeOff struct {
	//binding="required"
	//other validations gin provides: max, min, eq, ne, gt, lt, len, eqfield, nefield, dive, numeric, alpha
	//see https://github.com/go-playground/validator/tree/v8.18.1
	Type      string    `json:"reason" binding:"required"`
	Status    string    `json:"status" binding:"required"`
	Amount    float32   `json:"hours" binding:"required,gt=0"`
	StartDate time.Time `json:"startDate" binding:"required"`
}
