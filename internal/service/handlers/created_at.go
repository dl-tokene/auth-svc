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

func CreatedAt(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewCreatedAtRequest(r)
	if err != nil {
		logger.WithError(err).Debug("failed to parse request")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
		return
	}
	db := helpers.DB(r)

	userID, apiErr, err := helpers.Authenticate(helpers.AuthTypeSession, r)
	if apiErr != nil || err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, apiErr)
		return
	}

	userID, err = helpers.RetrieveRefreshToken(req.Data.SessionToken, r)
	if err != nil {
		logger.WithError(err).Debug("failed to retrieve session token")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
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

	result := resources.CreatedAt{
		Key:        resources.Key{Type: resources.USER, ID: fmt.Sprint(user.ID)},
		Attributes: resources.CreatedAtAttributes{CreatedAt: user.CreatedAt.String()},
	}

	ape.Render(w, result)
}
