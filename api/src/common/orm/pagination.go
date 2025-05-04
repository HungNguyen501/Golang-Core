package orm

type PaginationData[T any] struct {
	Total         int `json:"total"`
	CurrentOffset int `json:"current_offset"`
	Data          []T `json:"data"`
}
