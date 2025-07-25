package service

import "github.com/keshu12345/overlap/internal/model"

type OverlapService struct{}

func NewOverlapService() *OverlapService {
	return &OverlapService{}
}

func (s *OverlapService) CheckOverlap(r1, r2 model.DateRange) bool {
	return r1.Start.Before(r2.End) && r2.Start.Before(r1.End)
}
