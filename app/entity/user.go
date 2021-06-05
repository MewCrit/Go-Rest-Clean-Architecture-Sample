package entity

type User struct {
	Id        string `json:"id"`
	LoginType string `json:"loginType"`
	ImageUrl  string `json:"imageUrl"`
	FullName  string `json:"fullName"`
	IsActive  bool   `json:"isActive"`
	Base
}
