package handlers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewAdminLoginRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, apierrors.BadRequest(apierrors.CodeBadRequestData, err))
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
		ape.RenderErr(w, apierrors.BadRequest(apierrors.CodeAdminNotFound))
		return
	}

	// success logic
	doorman := helpers.DoormanConnector(r)
	pair, err := doorman.GenerateJwtPair(ethAddress, "session")
	if err != nil {
		logger.WithError(err).Error("failed to generate jwt")
		ape.RenderErr(w, apierrors.InternalError(err))
		return
	}

	ape.Render(w, pair)
}
