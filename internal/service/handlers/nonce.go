package handlers

import (
	"net/http"
	"time"

	nonces "github.com/LarryBattle/nonce-golang"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/util"
)

func GetNonce(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewNonceRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	address := request.Data.Attributes.Address
	termsHash := request.Data.Attributes.TermsHash
	db := helpers.DB(r)

	timeToExpire := helpers.ServiceConfig(r).NonceExpireTime
	expireTime := time.Now().UTC().Add(timeToExpire)
	nonceToken := nonces.NewToken()

	var message string = util.NonceToMessage(nonceToken)

	if termsHash != nil {
		message = util.NonceToTermsMessage(nonceToken, *termsHash)
	}

	nonce := data.Nonce{
		Message: nonceToken,
		Expires: &expireTime,
		Address: address,
	}

	// Required to make sure we have a clean `insert`-able state, as we're racing with nonce cleaner here
	err = db.Nonce().FilterByAddress(address).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	_, err = db.Nonce().Insert(nonce)
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewNonceModel(message))
}
