package service

import "go-clean-architecture/app/entity"

type ServiceInterce interface {
	GetUsers() []entity.User
}
