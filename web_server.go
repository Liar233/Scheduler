package main

import "net/http"

type WebServerAdapter struct {

}

func (wsa *WebServerAdapter) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("implement me")
}

func NewWebServer() *WebServerAdapter {
	return &WebServerAdapter{

	}
}
