package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/isshougai/rental-bookings/internal/config"
	"github.com/isshougai/rental-bookings/internal/handlers"
)

func routes(app *config.AppConfig) *gin.Engine {
	mux := gin.Default()
	mux.Use(sessions.Sessions("mysession", app.Store))

	nextHandler, wrapper := adapter.New()
	ns := NoSurf(nextHandler)
	mux.Use(wrapper(ns))

	mux.GET("/", handlers.Repo.Home)
	mux.GET("/about", handlers.Repo.About)
	mux.GET("/kiyomizu", handlers.Repo.Kiyomizu)
	mux.GET("/gion", handlers.Repo.Gion)

	mux.GET("/search-availability", handlers.Repo.Availability)
	mux.POST("/search-availability", handlers.Repo.PostAvailability)
	mux.POST("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.GET("/contact", handlers.Repo.Contact)

	mux.GET("/make-reservation", handlers.Repo.Reservation)
	mux.POST("/make-reservation", handlers.Repo.PostReservation)
	mux.GET("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Static("/static", "./static")

	return mux
}
