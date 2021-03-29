package pkg

import (
	"net/http"
)

type Application struct {
	Handler http.Handler
}

func NewApplication(handler http.Handler) Application {
	return Application{
		Handler: handler,
	}
}

