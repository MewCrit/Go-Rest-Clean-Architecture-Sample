package repository

import (
	"time"

	"go-clean-architecture/app/entity"
)

type repo struct{}

func RepoImpl() RepoInterface {
	return &repo{}
}

func (*repo) ReadUsers() []entity.User {

	users := make([]entity.User, 0, 6)

	users = append(users, entity.User{
		Id:        "547bf3d4-948d-48dc-b84e-eb7d5dd2dad9",
		LoginType: "facebook",
		ImageUrl:  "https://randomuser.me/api/portraits/women/48.jpg",
		FullName:  "Regina Garza",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	users = append(users, entity.User{
		Id:        "1b07521f-97b1-401f-bfed-1d14c27a87c9",
		LoginType: "google",
		ImageUrl:  "https://randomuser.me/api/portraits/women/13.jpg",
		FullName:  "Jen Carla",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	users = append(users, entity.User{
		Id:        "565c220f-0a65-42bd-aa03-9b1892639258",
		LoginType: "google",
		ImageUrl:  "https://randomuser.me/api/portraits/men/48.jpg",
		FullName:  "Swammi Swami",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	users = append(users, entity.User{
		Id:        "576caf7a-d8c8-4bd6-bc1e-bb0cd42427e1",
		LoginType: "google",
		ImageUrl:  "https://randomuser.me/api/portraits/men/18.jpg",
		FullName:  "Karl Sulliven",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	users = append(users, entity.User{
		Id:        "19593bd6-4287-4172-935b-a3512947e87f",
		LoginType: "google",
		ImageUrl:  "https://randomuser.me/api/portraits/men/11.jpg",
		FullName:  "Rodrigo Cortez",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	users = append(users, entity.User{
		Id:        "841a7f09-83e7-4951-9e39-efcc630c7bcc",
		LoginType: "google",
		ImageUrl:  "https://randomuser.me/api/portraits/women/11.jpg",
		FullName:  "Francine Chezka",
		IsActive:  false,
		Base: entity.Base{
			CreatedDate: time.Now(),
			UpdatedDate: time.Now()}})

	return users

}
