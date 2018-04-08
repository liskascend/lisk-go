package api

import "strings"

type (
	// Error represents an API error
	Error struct {
		// Message is the root error message
		Message string `json:"message"`
		// Errors contains detailed errors
		Errors []struct {
			// Code is the error code
			Code string `json:"code"`
			// Name is the name of the error
			Name string `json:"name"`
			// In is the part of the request that contains the error (e.g. body)
			In string `json:"in"`
			// Message is the detail error message
			Message string `json:"message"`
			// Errors contains deep errors
			Errors []struct {
				// Code is the error code
				Code string `json:"code"`
				// Message is the detail error message
				Message string `json:"message"`
				// Description is a human readable error description
				Description string `json:"description"`
				// Path are the body nodes which contain the error
				Path []string `json:"path"`
			} `json:"errors"`
		} `json:"errors"`
	}
)

// Error prints API errors in a human readable format
func (e *Error) Error() string {
	b := strings.Builder{}

	if len(e.Errors) > 0 {
		b.WriteString(e.Message + " :\n")
	} else {
		b.WriteString(e.Message + "\n")
	}

	for _, detailError := range e.Errors {
		b.WriteString(detailError.Message + ":\n")

		for _, deepError := range detailError.Errors {
			b.WriteString("\t" + strings.Join(deepError.Path, ",") + ": " + deepError.Message + "\n")
		}
	}
	return b.String()
}
