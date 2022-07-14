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

	mux.GET("/", gin.HandlerFunc(handlers.Repo.Home))
	mux.GET("/about", gin.HandlerFunc(handlers.Repo.About))
	mux.GET("/kiyomizu", gin.HandlerFunc(handlers.Repo.Kiyomizu))
	mux.GET("/gion", gin.HandlerFunc(handlers.Repo.Gion))

	mux.GET("/search-availability", gin.HandlerFunc(handlers.Repo.Availability))
	mux.POST("/search-availability", gin.HandlerFunc(handlers.Repo.PostAvailability))
	mux.POST("/search-availability-json", gin.HandlerFunc(handlers.Repo.AvailabilityJSON))

	mux.GET("/contact", gin.HandlerFunc(handlers.Repo.Contact))

	mux.GET("/make-reservation", gin.HandlerFunc(handlers.Repo.Reservation))
	mux.POST("/make-reservation", gin.HandlerFunc(handlers.Repo.PostReservation))

	mux.Static("/static", "./static")

	return mux
}
