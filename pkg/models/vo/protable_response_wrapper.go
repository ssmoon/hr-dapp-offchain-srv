package vo

type ProTableResponseWrapper[T any] struct {
	Data    []T   `json:"data"`
	Success bool  `default:"true" json:"success"`
	Total   int64 `json:"total"`
}
