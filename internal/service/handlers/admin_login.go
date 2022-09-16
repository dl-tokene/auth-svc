package handlers

import (
	"errors"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
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
	db := helpers.DB(r)
	nodeAdmins := helpers.NodeAdmins(r)
	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

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

	if !nodeAdmins.CheckAdmin(common.HexToAddress(ethAddress)) {
		logger.Debug("not admin's address")
		ape.RenderErr(w, problems.BadRequest(errors.New("not admin's address"))...)
		return
	}
	// success logic
	token, err := helpers.GenerateAdminJWT(ethAddress, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Error("failed to generate a token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	refreshToken, err := helpers.GenerateAdminRefreshToken(ethAddress, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Error("failed to generate a token")
		ape.Render(w, problems.InternalError())
		return
	}

	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewAdminLoginModel(
		token,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime))
}
