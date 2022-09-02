package handlers

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	errors "gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewAdminLoginRequest(r)
	if err != nil {
		logger.WithError(err).Debug("bad request")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
		return
	}
	db := helpers.DB(r)
	nodeAdmins := helpers.NodeAdmins(r)
	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

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

	if !nodeAdmins.CheckAdmin(common.HexToAddress(ethAddress)) {
		logger.Debug("not admin's address")
		ape.RenderErr(w, errors.BadRequest(errors.CodeUnauthorized))
		return
	}
	// success logic
	token, err := helpers.GenerateAdminJWT(ethAddress, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a token"
		logger.WithError(err).Error(details)
		ape.RenderErr(w, errors.InternalError(details))
		return
	}
	refreshToken, err := helpers.GenerateAdminRefreshToken(ethAddress, helpers.ServiceConfig(r))
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

	ape.Render(w, models.NewAdminLoginModel(
		token,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime))
}
