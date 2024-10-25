package web

type CategoryUpdateRequest struct {
	Id int 			`validate:"required" json:"id"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}