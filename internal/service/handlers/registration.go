package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
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
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// success logic
	var user *data.User
	user, err = db.Users().Insert(data.User{Address: ethAddress})
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	token, err := helpers.GenerateJWT(user, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Error("failed to generate a token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(user, helpers.ServiceConfig(r))
	if err != nil {
		logger.WithError(err).Error("failed to generate a refresh token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	err = db.Nonce().FilterByAddress(ethAddress).Delete()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	result := models.NewRegistrationModel(
		token,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime,
		*user,
	)

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, result)
}
