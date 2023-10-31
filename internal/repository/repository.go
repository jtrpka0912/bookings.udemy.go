package repository

import "github.com/jtrpka0912/bookings.udemy.go/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}