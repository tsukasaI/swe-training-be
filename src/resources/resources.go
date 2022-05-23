package resources

type CommonResponseField struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PostResponse struct {
	CommonResponseField
	Comment string `json:"comment"`
	User    UserResponse
}

type UserResponse struct {
	CommonResponseField
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Posts   *[]PostResponse `json:"posts"`
	Follows *[]UserResponse `json:"follows"`
}
