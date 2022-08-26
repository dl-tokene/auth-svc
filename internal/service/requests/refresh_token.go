package requests

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func NewRefreshToken(r *http.Request) (resources.RefreshTokenRequest, error) {
	request := resources.RefreshTokenRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}
	return request, nil
}
