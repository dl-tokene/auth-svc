package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	db := helpers.DB(r)
	doorman := helpers.DoormanConnector(r)

	ethAddress, refreshToken, err := helpers.ValidateJWT(doorman, r)
	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, apierrors.Unauthorized(apierrors.CodeUnauthorized))
		return
	}

	pair, err := doorman.RefreshJwt(refreshToken)
	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, apierrors.Unauthorized(apierrors.CodeUnauthorized))
		return
	}

	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, apierrors.InternalError(err))
		return
	}

	ape.Render(w, pair)
}
