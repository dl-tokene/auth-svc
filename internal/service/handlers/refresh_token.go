package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	db := helpers.DB(r)
	doorman := helpers.DoormanConnector(r)

	refreshToken, err := doorman.GetAuthToken(r)
	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	ethAddress, err := doorman.ValidateJwt(refreshToken)
	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}
	pair, err := doorman.RefreshJwt(refreshToken)

	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, pair)
}
