package api

import "strings"

type (
	Error struct {
		Message string `json:"message"`
		Errors []struct {
			Code    string `json:"code"`
			Name    string `json:"name"`
			In      string `json:"in"`
			Message string `json:"message"`
			Errors []struct {
				Code        string   `json:"code"`
				Message     string   `json:"message"`
				Description string   `json:"description"`
				Path        []string `json:"path"`
			} `json:"errors"`
		} `json:"errors"`
	}
)

func (e *Error) Error() string {
	b := strings.Builder{}
	b.WriteString(e.Message + " :\n")

	for _, detailError := range e.Errors {
		b.WriteString(detailError.Message + ":\n")

		for _, deepError := range detailError.Errors {
			b.WriteString("\t" + deepError.Message + "\n")
		}
	}
	return b.String()
}
