package requests

import (
	"encoding/json"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/util"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

type NonceRequest struct {
	Data resources.AuthNonceRequest
}

func NewNonceRequest(r *http.Request) (NonceRequest, error) {
	var request NonceRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	request.Data.Attributes.Address = strings.ToLower(request.Data.Attributes.Address)
	return request, request.validate()
}

func (r *NonceRequest) validate() error {
	return validation.Errors{
		"/data/type": validation.Validate(&r.Data.Type, validation.Required, validation.In(resources.AUTH_NONCE_REQUEST)),
		"/data/attributes/address": validation.Validate(&r.Data.Attributes.Address,
			validation.Required,
			validation.Match(util.AddressRegexp)),
	}.Filter()
}
