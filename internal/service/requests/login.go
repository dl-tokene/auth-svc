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

type LoginRequest struct {
	Data resources.LoginRequest
}

func NewLoginRequest(r *http.Request) (LoginRequest, error) {
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	request.Data.Attributes.AuthPair.Address = strings.ToLower(request.Data.Attributes.AuthPair.Address)
	return request, request.validate()
}

func (r *LoginRequest) validate() error {
	return validation.Errors{
		"/data/type": validation.Validate(&r.Data.Type, validation.Required, validation.In(resources.LOGIN_REQUEST)),
		"/data/attributes/auth_pair/address": validation.Validate(&r.Data.Attributes.AuthPair.Address,
			validation.Required,
			validation.Match(util.AddressRegexp)),
		"/data/attributes/auth_pair/signed_message": validation.Validate(&r.Data.Attributes.AuthPair.SignedMessage,
			validation.Required,
			validation.Match(util.SignatureRegexp)),
	}.Filter()
}
