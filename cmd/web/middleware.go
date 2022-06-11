package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justinas/nosurf"
)

func WriteToConsole() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hit the page")
		c.Next()
	}
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
