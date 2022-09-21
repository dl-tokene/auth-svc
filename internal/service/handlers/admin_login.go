package handlers

import (
	"errors"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewAdminLoginRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nodeAdmins := helpers.NodeAdmins(r)
	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

	apiErr, err := helpers.ValidateNonce(ethAddress, signature, r)
	if err != nil {
		logger.WithError(err).Debug("failed to validate nonce")
		ape.RenderErr(w, apiErr)
		return
	}

	if !nodeAdmins.CheckAdmin(common.HexToAddress(ethAddress)) {
		logger.Debug("not admin's address")
		ape.RenderErr(w, problems.BadRequest(errors.New("not admin's address"))...)
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
