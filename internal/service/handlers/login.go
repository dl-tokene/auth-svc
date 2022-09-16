package handlers

import (
	"errors"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
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
	db := helpers.DB(r)

	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

	address, err := db.Users().FilterByAddress(ethAddress).Get()
	if address == nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	nonce, err := db.Nonce().FilterByAddress(ethAddress).Get()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nonce == nil {
		logger.WithField("address", ethAddress).Debug("nonce not found on login")
		ape.RenderErr(w, problems.BadRequest(errors.New("nonce does not exist"))...)
		return
	}

	err = helpers.VerifySignature(helpers.NonceToHash(nonce), signature, ethAddress)
	if err != nil {
		logger.WithError(err).Debug("signature verification failed")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// success logic
	token, err := helpers.GenerateJWT(address, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Error("failed to generate a token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(address, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Error("failed to generate a refresh token")
		ape.Render(w, problems.InternalError())
		return
	}

	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewLoginModel(
		token,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime))
}
