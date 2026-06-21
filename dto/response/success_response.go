package response

type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Result  T      `json:"result"`
}
