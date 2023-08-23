package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Success(w *http.ResponseWriter, status int, data any) {
	meta := Meta{
		Success: true,
		Message: "success",
	}
	res := Response{
		Meta: meta,
		Data: data,
	}

	json.NewEncoder(*w).Encode(res)
}

func Fail(w *http.ResponseWriter, status int, errorMessage string) {
	meta := Meta{
		Success: false,
		Message: errorMessage,
	}
	res := Response{
		Meta: meta,
		Data: nil,
	}

	json.NewEncoder(*w).Encode(res)
}
