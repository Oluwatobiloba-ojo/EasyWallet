package response

type ApiResponse struct {
	Data   any
	Status bool
}

func NewApiResponse[T any](data T, status bool) *ApiResponse {
	return &ApiResponse{Data: data, Status: status}
}
