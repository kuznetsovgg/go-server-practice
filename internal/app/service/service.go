package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CountDays() int64 {
	d := time.Date(2026, 01, 01, 0, 0, 0, 0, time.Local)
	dur := time.Until(d)
	return int64(dur.Hours()) / 24
}
