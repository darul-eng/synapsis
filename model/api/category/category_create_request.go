package category

type CategoryCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
