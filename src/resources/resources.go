package resources

type responseBody struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type commonModelResponseField struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type PostResponse struct {
	commonModelResponseField
	Comment string              `json:"comment"`
	User    UserResponseForHome `json:"writer"`
}

type UserResponseForHome struct {
	commonModelResponseField
	Name string `json:"name"`
}

func CreateResponseBody(code string, data interface{}) responseBody {
	return responseBody{
		Code: code,
		Data: data,
	}
}
