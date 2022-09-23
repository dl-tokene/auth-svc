package handlers

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	errors "gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	db := helpers.DB(r)

	userID, token, apiErr, err := helpers.Authenticate(helpers.AuthTypeSession, r)
	if apiErr != nil || err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, apiErr)
		return
	}

	userID, err = helpers.RetrieveRefreshToken(token, r)
	if err != nil {
		logger.WithError(err).Debug("failed to retrieve refresh token")
		ape.RenderErr(w, errors.Unauthorized(errors.CodeUnauthorized, err))
		return
	}

	user, err := db.Users().FilterByUserID(userID).Get()
	if err != nil {
		logger.WithError(err).Error("failed to query db")
		ape.RenderErr(w, errors.InternalError(errors.InternalError(), err))
		return
	}
	if user == nil {
		ape.RenderErr(w, errors.NotFound(fmt.Sprintf("User with id %d doesn't exist", userID)))
		return
	}

	// success logic
	sessionToken, err := helpers.GenerateJWT(user, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a token"
		logger.WithError(err).Error(details)
		ape.RenderErr(w, errors.InternalError(details))
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(user, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a refresh token"
		logger.WithError(err).Error(details)
		ape.Render(w, errors.InternalError(details))
		return
	}

	ape.Render(w, models.NewLoginModel(
		sessionToken,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime))
}
