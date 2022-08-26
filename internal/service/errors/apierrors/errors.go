package apierrors

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io"
	"net/http"
)

// TODO(rufus): cover all 4XX errors

func BadRequest(code string, extra ...interface{}) *jsonapi.ErrorObject {
	if code == "" {
		code = CodeBadRequest
	}
	result := baseError(code, http.StatusBadRequest)

	for _, arg := range extra {
		if arg == nil {
			continue
		}
		switch arg := arg.(type) {
		case problems.BadRequester:
			result.Meta = mapValidationErrors(arg.BadRequest())
		case validation.Errors:
			result.Meta = mapValidationErrors(arg)
		case error:
			if errors.Cause(arg) == io.EOF {
				result.Code = CodeEmptyBody
				result.Detail = detailsMap[CodeEmptyBody]
			} else {
				result.Meta = &map[string]interface{}{"error": arg.Error()}
			}
		case string:
			addDetails(&result, arg)
		default:
			addDetails(&result, fmt.Sprint(arg))
		}
	}

	return &result
}

func Unauthorized(code string, extra ...interface{}) *jsonapi.ErrorObject {
	if code == "" {
		code = CodeUnauthorized
	}
	result := baseError(code, http.StatusUnauthorized)
	processExtraDetailsDefault(&result, extra...)
	return &result
}

func Forbidden(code string, extra ...interface{}) *jsonapi.ErrorObject {
	if code == "" {
		code = CodeForbidden
	}
	result := baseError(code, http.StatusForbidden)
	processExtraDetailsDefault(&result, extra...)
	return &result
}

func NotFound(extra ...interface{}) *jsonapi.ErrorObject {
	result := baseError(CodeNotFound, http.StatusNotFound)
	processExtraDetailsDefault(&result, extra...)
	return &result
}

func Conflict(extra ...interface{}) *jsonapi.ErrorObject {
	result := baseError(CodeConflict, http.StatusConflict)
	processExtraDetailsDefault(&result, extra...)
	return &result
}

func TooEarly(extra ...interface{}) *jsonapi.ErrorObject {
	result := baseError(CodeTooEarly, http.StatusTooEarly)
	processExtraDetailsDefault(&result, extra...)
	return &result
}

func InternalError(extra ...interface{}) *jsonapi.ErrorObject {
	result := baseError(CodeInternalError, http.StatusInternalServerError)
	processExtraDetailsDefault(&result, extra...)
	return &result
}
