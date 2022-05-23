package controllers

type CommonResponseField struct {
	Id        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
