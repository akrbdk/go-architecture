package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() (int, int64) {
	nextYear := time.Now().Year() + 1
	durToNewYear := time.Until(time.Date(nextYear, time.January, 1, 0, 0, 0, 0, time.UTC))
	return nextYear, int64(durToNewYear.Hours()) / 24
}
