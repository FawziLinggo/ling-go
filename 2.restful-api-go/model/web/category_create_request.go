package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=200,min=0" json:"name"`
}
