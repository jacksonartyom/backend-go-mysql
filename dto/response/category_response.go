package response

type CategoryResponse struct {
	CategoryId string `json:"_id"`
	Name       string
	Type       string
}
