package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/isshougai/rental-bookings/pkg/config"
	"github.com/isshougai/rental-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) *gin.Engine {
	mux := gin.Default()
	mux.Use(sessions.Sessions("mysession", app.Store))

	nextHandler, wrapper := adapter.New()
	ns := NoSurf(nextHandler)
	mux.Use(wrapper(ns))

	mux.GET("/", gin.HandlerFunc(handlers.Repo.Home))
	mux.GET("/about", gin.HandlerFunc(handlers.Repo.About))

	mux.Static("/static/images", "./static/images")

	return mux
}