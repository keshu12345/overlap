package model

import "time"

type DateRange struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}

type OverlapRequest struct {
	Range1 DateRange `json:"range1" binding:"required"`
	Range2 DateRange `json:"range2" binding:"required"`
}

type OverlapResponse struct {
	Overlap bool `json:"overlap"`
}
