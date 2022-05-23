package controllers

type UserResponse struct {
	CommonResponseField
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Posts   *[]PostResponse `json:"posts"`
	Follows *[]UserResponse `json:"follows"`
}
