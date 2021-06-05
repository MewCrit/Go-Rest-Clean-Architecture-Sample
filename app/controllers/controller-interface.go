package controller

import (
	"net/http"
)

type ControllerInterface interface {
	GetUsers(response http.ResponseWriter, request *http.Request)
}
