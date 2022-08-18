package handlers

import (
	"fmt"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	errors "gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewRefreshToken(r)
	if err != nil {
		logger.WithError(err).Info("failed to parse request")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
		return
	}
	db := helpers.DB(r)

	userID, apiErr, err := helpers.Authenticate(helpers.AuthTypeSession, r)
	if apiErr != nil || err != nil {
		logger.WithError(err).Info("failed authentication")
		ape.RenderErr(w, apiErr)
		return
	}

	userID, err = helpers.RetrieveRefreshToken(request.RefreshToken, r)
	if err != nil {
		logger.WithError(err).Info("failed to retrieve refresh token")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
		return
	}

	user := db.Users().FilterByUserID(userID).Get()
	if user == nil {
		ape.RenderErr(w, errors.NotFound(fmt.Sprintf("User with id %d doesn't exist", userID)))
		return
	}
	address := db.Users().FilterByUserID(userID).Get()
	if address == nil {
		errorText := fmt.Sprintf("address doesn't exist for user with ID %d", userID)
		logger.Error(errorText)
		ape.RenderErr(w, errors.InternalError(errorText))
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

	result := resources.RefreshTokenResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}
	ape.Render(w, result)
}
