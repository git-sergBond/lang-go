package service

import (
	"net/http"
	"strconv"
)

// GifExampleHandlerGetParameters
// parse URL parameters:
// - cycles
func GifExampleHandlerGetParameters(request *http.Request) map[string]any {
	if err := request.ParseForm(); err != nil {
		FprintfError("request.ParseForm()", request, err)
	}

	result := map[string]any{}

	if cycles, err := strconv.Atoi(request.Form["cycles"][0]); err == nil {
		result["cycles"] = cycles
	}

	return result
}
