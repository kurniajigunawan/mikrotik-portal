package service

import "time"

type status int8

const (
	StatusInactive status = iota
	StatusActive
)

type Service struct {
	ID           int64
	Name         string
	ClientID     string
	ClientSecret string
	Status       status
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
