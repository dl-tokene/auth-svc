package handlers

import (
	"errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/resources"
	"net/http"
)

func CreatedAt(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	db := helpers.DB(r)
	doorman := helpers.DoormanConnector(r)

	address, _, _, err := helpers.ValidatePurposeJWT(doorman, r)
	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, apierrors.Unauthorized(apierrors.CodeUnauthorized))
		return
	}

	user, err := db.Users().FilterByAddress(address).Get()
	if user == nil {
		ape.RenderErr(w, apierrors.NotFound(errors.New("User not found")))
		return
	}

	// success logic

	result := resources.CreatedAtResponse{
		Data: resources.CreatedAt{
			Key: resources.Key{Type: resources.CREATED_AT},
			Attributes: resources.CreatedAtAttributes{
				Address:   address,
				CreatedAt: user.CreatedAt,
			},
		},
	}

	ape.Render(w, result)
}
