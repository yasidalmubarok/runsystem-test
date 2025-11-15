package dto

type UserRequest struct {
	Name  string `json:"name" binding:"required"`
	Hobbies string `json:"hobbies" binding:"required"`
}