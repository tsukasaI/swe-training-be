package controllers

type PostResponse struct {
	CommonResponseField
	Comment string `json:"comment"`
	User    UserResponse
}
