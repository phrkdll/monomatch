package utils

import (
	"encoding/json"
	"net/http"

	"github.com/phrkdll/must/pkg/must"
)

func SendJSON[T any](w http.ResponseWriter, status int, data T) {
	typedResult := must.NewTypedResult(must.UntypedResult{}, data)
	jsonData := must.Return(json.Marshal(&typedResult)).
		ElseRespond(w, http.StatusBadRequest)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
