package api

type (
	SortMode string

	ListOptions struct {
		Limit  int
		Offset int
		Sort   SortMode
	}

	GenericResponse struct {
		Meta  interface{} `json:"meta"`
		Links interface{} `json:"links"`
	}

	LiskAmount int64
)

const (
	SortModeAscending  SortMode = "ASC"
	SortModeDescending SortMode = "ASC"
)
