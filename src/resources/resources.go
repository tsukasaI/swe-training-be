package resources

type ResponseBody struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type CommonModelResponseField struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PostResponse struct {
	CommonModelResponseField
	Comment string       `json:"comment"`
	User    UserResponse `json:"writer"`
}

type UserResponse struct {
	CommonModelResponseField
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Posts   *[]PostResponse `json:"posts"`
	Follows *[]UserResponse `json:"follows"`
}

func CreateResponseBody(code string, data interface{}) ResponseBody {
	return ResponseBody{
		Code: code,
		Data: data,
	}
}
