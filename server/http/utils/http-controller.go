package utils

import "github.com/gorilla/mux"

type HTTPController interface {
	HandleRoutes(*mux.Router)
}
