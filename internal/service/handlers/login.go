package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func Login(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewLoginRequest(r)
	if err != nil {

		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

	apiErr, err := helpers.ValidateNonce(ethAddress, signature, r)
	if err != nil {
		logger.WithError(err).Debug("failed to validate nonce")
		ape.RenderErr(w, apiErr)
		return
	}

	// success logic
	doorman := helpers.DoormanConnector(r)
	pair, err := doorman.GenerateJwtPair(ethAddress, "session")
	if err != nil {
		logger.WithError(err).Error("failed to generate jwt")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, pair)
}
