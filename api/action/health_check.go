package action

import (
	"net/http"
	"fmt"
)

type HealthCheck struct {
}

func (hc *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}
