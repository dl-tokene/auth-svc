package apierrors

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
)

func baseError(code string, statusCode int) jsonapi.ErrorObject {
	details, ok := detailsMap[code]
	if !ok {
		details = ""
	}
	return jsonapi.ErrorObject{
		Title:  http.StatusText(statusCode),
		Detail: details,
		Status: fmt.Sprintf("%d", statusCode),
		Code:   code,
	}
}

func mapValidationErrors(errors map[string]error) *map[string]interface{} {
	result := make(map[string]interface{}, len(errors))
	for k, v := range errors {
		result[k] = v.Error()
	}
	return &result
}

func addDetails(err *jsonapi.ErrorObject, details string) {
	if err.Detail != "" {
		err.Detail = fmt.Sprintf("%v; %v", err.Detail, details)
	} else {
		err.Detail = details
	}
}

func processExtraDetailsDefault(err *jsonapi.ErrorObject, extra ...interface{}) {
	if len(extra) > 0 && extra[0] != nil {
		err.Detail = ""
	}
	for _, arg := range extra {
		if arg == nil {
			continue
		}
		switch arg := arg.(type) {
		case error:
			err.Meta = &map[string]interface{}{"error": arg.Error()}
		case string:
			addDetails(err, arg)
		default:
			addDetails(err, fmt.Sprint(arg))
		}
	}
}

func Details(code string) (details string) {
	return detailsMap[code]
}
