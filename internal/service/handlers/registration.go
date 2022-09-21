package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func Register(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewRegistrationRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	db := helpers.DB(r)

	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

	apiErr, err := helpers.ValidateNonce(ethAddress, signature, r)
	if err != nil {
		logger.WithError(err).Debug("failed to validate nonce")
		ape.RenderErr(w, apiErr)
		return
	}

	existingAddress, err := db.Users().FilterByAddress(ethAddress).Get()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if existingAddress != nil {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	// success logic
	_, err = db.Users().Insert(data.User{Address: ethAddress, CreatedAt: time.Now()})
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	doorman := helpers.DoormanConnector(r)
	pair, err := doorman.GenerateJwtPair(ethAddress, "session")
	if err != nil {
		logger.WithError(err).Error("failed to generate jwt")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, pair)
}
