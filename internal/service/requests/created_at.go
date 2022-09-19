package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

type CreatedAtRequest struct {
	Data resources.SessionTokenRequest
}

func NewCreatedAtRequest(r *http.Request) (CreatedAtRequest, error) {
	var request CreatedAtRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, request.validate()
}

func (r *CreatedAtRequest) validate() error {
	return validation.Errors{
		"/data/session_token": validation.Validate(&r.Data.SessionToken, validation.Required),
	}.Filter()
}
