package api

import (
	"github.com/braintree/manners"
	"net/http"
	"github.com/Liar233/Scheduler/config"
	"fmt"
)

type WebServerAdapter struct {
	server *manners.GracefulServer
}

// start serving http
func (wsa *WebServerAdapter) ListenAndServe() {
	wsa.server.ListenAndServe()
}

// Stop all api routes handlers
func (wsa *WebServerAdapter) Close() {
	wsa.server.Close()
}

func NewWebServer(router *RouterAdapter, conf *config.AppConfig) *WebServerAdapter {
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%d", conf.Port),
	}

	return &WebServerAdapter{
		server: manners.NewWithServer(server),
	}
}
