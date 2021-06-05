package repository

import (
	"go-clean-architecture/app/entity"
)

type RepoInterface interface {
	ReadUsers() []entity.User
}
