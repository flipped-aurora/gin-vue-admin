//go:build windows
// +build windows

package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    30 * time.Second, // increasing timeout from 20 sec to 30 sec by mohamed hassan
		WriteTimeout:   30 * time.Second, // increasing timeout from 20 sec to 30 sec by mohamed hassan
		MaxHeaderBytes: 1 << 20,
	}
}
