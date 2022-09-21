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

	// validate request
	nonce, err := db.Nonce().FilterByAddress(ethAddress).Get()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nonce == nil {
		logger.WithField("address", ethAddress).Debug("nonce not found")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	err = helpers.VerifySignature(helpers.NonceToHash(nonce), signature, ethAddress)
	if err != nil {
		logger.WithError(err).Debug("signature verification failed")
		ape.RenderErr(w, problems.Unauthorized())
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
	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, pair)
}
