package main

import (
	"net/http"
	"fmt"
	"time"
	"github.com/braintree/manners"
	"sync"
)

type WebServerAdapter struct {
	server *manners.GracefulServer
	wg     *sync.WaitGroup
}

// start serving http
func (wsa *WebServerAdapter) ListenAndServe() error {
	return wsa.server.ListenAndServe()
}

// Stop all api routes handlers
func (wsa *WebServerAdapter) Close() {
	wsa.server.Close()
	wsa.wg.Done()
}

// Create new wrapped web-server
func NewWebServer(config *AppConfig, wg *sync.WaitGroup) WebServerAdapter {
	controller := NewApiController()

	httpTimeout := time.Duration(config.ApiTimeout) * time.Second

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		ReadTimeout:  httpTimeout,
		WriteTimeout: httpTimeout,
		Handler:      controller.router,
	}

	server := manners.NewWithServer(&httpServer)

	return WebServerAdapter{
		server: server,
		wg:     wg,
	}
}
