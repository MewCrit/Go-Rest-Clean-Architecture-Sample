package service

import (
	"go-clean-architecture/app/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) ReadUsers() []entity.User {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.User)
}

func TestGetUsers(t *testing.T) {

	mockRepo := new(MockRepository)

	users := make([]entity.User, 0, 1)

	users = append(users, entity.User{
		Id:        "a",
		LoginType: "a",
		ImageUrl:  "a",
		FullName:  "a",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	mockRepo.On("GetUsers").Return(users)
	testService := ServiceImpl(mockRepo)

	result := testService.GetUsers()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "test", result[0].Id)
	assert.Equal(t, "fb", result[0].LoginType)

}
