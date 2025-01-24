package response

type (
	BaseResponse[T any] struct {
		StatusCode int    `json:"status_code"`
		Success    bool   `json:"success"`
		Message    string `json:"message"`
		Data       T      `json:"data"`
	}

	IBaseResponse[T any] interface {
		IsSuccess() bool
		GetMessage() string
		GetData() T
	}
)

func (b *BaseResponse[T]) IsSuccess() bool {
	return b.Success
}

func (b *BaseResponse[T]) GetMessage() string {
	return b.Message
}

func (b *BaseResponse[T]) GetData() T {
	return b.Data
}
