package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	errors "gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func Login(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewLoginRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
		return
	}
	db := helpers.DB(r)

	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

	address, err := db.Users().FilterByAddress(ethAddress).Get()
	if address == nil {
		ape.RenderErr(w, errors.BadRequest(errors.CodeNotFound))
		return
	}
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, errors.InternalError(errors.InternalError(), err))
		return
	}
	nonce, err := db.Nonce().FilterByAddress(ethAddress).Get()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, errors.InternalError(errors.InternalError(), err))
		return
	}
	if nonce == nil {
		logger.WithField("address", ethAddress).Debug("nonce not found on login")
		ape.RenderErr(w, errors.BadRequest(errors.CodeNonceNotFound))
		return
	}

	err = helpers.VerifySignature(helpers.NonceToHash(nonce), signature, ethAddress)
	if err != nil {
		logger.WithError(err).Debug("signature verification failed")
		ape.RenderErr(w, errors.BadRequest(errors.CodeSignatureVerificationFailed, err))
		return
	}

	// success logic
	token, err := helpers.GenerateJWT(address, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a token"
		logger.WithError(err).Error(details)
		ape.RenderErr(w, errors.InternalError(details))
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(address, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a refresh token"
		logger.WithError(err).Error(details)
		ape.Render(w, errors.InternalError(details))
		return
	}

	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, errors.InternalError(errors.InternalError(), err))
		return
	}

	ape.Render(w, models.NewLoginModel(
		token,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime))
}