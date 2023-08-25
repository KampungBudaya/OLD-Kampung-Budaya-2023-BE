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
	Message string `json:"message"`
}

func Success(w *http.ResponseWriter, status int, data interface{}) {
	meta := Meta{
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
		Message: errorMessage,
	}
	res := Response{
		Meta: meta,
		Data: nil,
	}

	json.NewEncoder(*w).Encode(res)
}
