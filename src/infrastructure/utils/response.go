package utils

type IResponse[TData, ITMeta any] interface {
	MakeMetaData(message string, data map[string]ITMeta, page *int, perPage *int, totalPage *int)
	BuildResponse() *response[TData, ITMeta]
}

type Meta[T any] struct {
	Message    string       `json:"message,omitempty"`
	Additional map[string]T `json:"additional,omitempty"`
	Page       *int         `json:"page,omitempty"`
	PerPage    *int         `json:"per_page,omitempty"`
	TotalPage  *int         `json:"total_page,omitempty"`
}

type response[TData, TMeta any] struct {
	Data TData        `json:"data"`
	Meta *Meta[TMeta] `json:"meta,omitempty"`
}

func (r *response[TData, TMeta]) BuildResponse() *response[TData, TMeta] {
	return r
}

func (r *response[TData, ITMeta]) MakeMetaData(message string, data map[string]ITMeta, page *int, perPage *int, totalPage *int) {
	r.Meta = &Meta[ITMeta]{
		Message:    message,
		Page:       page,
		PerPage:    perPage,
		TotalPage:  totalPage,
		Additional: data,
	}
}

func NewResponse[T, TMeta any](data T) IResponse[T, TMeta] {
	return &response[T, TMeta]{
		Data: data,
	}
}
