package service

import (
	"go-clean-architecture/app/entity"
	"go-clean-architecture/app/repository"
)

type service struct{}

var (
	repo repository.RepoInterface
)

func ServiceImpl(repository repository.RepoInterface) ServiceInterce {
	repo = repository
	return &service{}
}

func (*service) GetUsers() []entity.User {
	return repo.ReadUsers()
}
