package helpers

import (
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/doorman/connector"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
)

func GetNonce(address string, r *http.Request) (*data.Nonce, *jsonapi.ErrorObject, error) {
	db := DB(r)
	nonce, err := db.Nonce().FilterByAddress(address).Get()
	if err != nil {
		return nil, problems.InternalError(), errors.New("failed to query db")
	}

	if nonce == nil {
		return nil, apierrors.BadRequest(apierrors.CodeNonceNotFound), errors.New("nonce not found")
	}

	err = db.Nonce().FilterByAddress(address).Delete()
	if err != nil {
		return nil, apierrors.InternalError(err), errors.New("failed to query db")
	}
	return nonce, nil, nil
}
func ValidateNonce(address string, signature string, r *http.Request) (*jsonapi.ErrorObject, error) {
	nonce, apierr, err := GetNonce(address, r)
	if err != nil {
		return apierr, err
	}
	err = VerifySignature(NonceToHash(nonce), signature, address)
	if err != nil {
		return apierrors.BadRequest(apierrors.CodeSignatureVerificationFailed, err), err
	}
	return nil, nil
}
func ValidateJWT(doorman connector.ConnectorI, r *http.Request) (string, string, error) {

	token, err := doorman.GetAuthToken(r)
	if err != nil {
		return "", "", err
	}
	address, err := doorman.ValidateJwt(token)
	if err != nil {
		return "", "", err
	}
	return address, token, err
}
func ValidatePurposeJWT(doorman connector.ConnectorI, r *http.Request) (string, string, string, error) {
	address, token, err := ValidateJWT(doorman, r)
	if err != nil {
		return "", "", "", err
	}

	purpose, err := doorman.CheckPurpose(token)
	if err != nil {
		return "", "", "", err
	}

	return address, token, purpose, nil
}
