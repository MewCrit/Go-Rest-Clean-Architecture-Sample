package controller

import (
	"encoding/json"
	"net/http"

	"go-clean-architecture/app/service"
)

type controller struct{}

var (
	serviceImpl service.ServiceInterce
)

func Controller(service service.ServiceInterce) ControllerInterface {
	serviceImpl = service
	return &controller{}
}

func (*controller) GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	users := serviceImpl.GetUsers()
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}
