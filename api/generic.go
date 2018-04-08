package api

type (
	// SortMode specifies how results are sorted
	SortMode string

	// ListOptions are options which are used for pagination
	ListOptions struct {
		// Limit is the amount of results requested
		Limit int
		// Offset is the offset from the first item from which on results should be returned
		Offset int
		// Sort is the mode in which results should be sorted
		Sort SortMode
	}

	// GenericResponse specifies the generic meta fields of an API response
	GenericResponse struct {
		// Meta is meta information of the response
		Meta interface{} `json:"meta"`
		// Links is a currently unused field
		Links interface{} `json:"links"`
	}

	// LiskAmount is an amount of Lisk
	LiskAmount int64
)

const (
	// SortModeAscending is the raw SortMode for ascending sorting
	SortModeAscending SortMode = "ASC"
	// SortModeDescending is the raw SortMode for descending sorting
	SortModeDescending SortMode = "DESC"
)
