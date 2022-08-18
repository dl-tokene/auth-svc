package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
	errors "gitlab.com/tokene/nonce-auth-svc/internal/service/errors/apierrors"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/models"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/requests"
)

func Register(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewRegistrationRequest(r)
	if err != nil {
		logger.WithError(err).Info("bad request")
		ape.RenderErr(w, errors.BadRequest(errors.CodeBadRequestData, err))
		return
	}
	db := helpers.DB(r)

	ethAddress := request.Data.Attributes.AuthPair.Address
	signature := request.Data.Attributes.AuthPair.SignedMessage

	existingAddress := db.Users().FilterByAddress(ethAddress).Get()
	if existingAddress != nil {
		ape.RenderErr(w, errors.Conflict(errors.Details(errors.CodeAddressExists)))
		return
	}

	// validate request
	nonce := db.Nonce().FilterByAddress(ethAddress).Get()
	if nonce == nil {
		logger.WithField("address", ethAddress).Info("nonce not found")
		ape.RenderErr(w, errors.BadRequest(errors.CodeNonceNotFound))
		return
	}

	err = helpers.VerifySignature(helpers.NonceToHash(nonce), signature, ethAddress)
	if err != nil {
		logger.WithError(err).Info("signature verification failed")
		ape.RenderErr(w, errors.BadRequest(errors.CodeSignatureVerificationFailed, err))
		return
	}

	// success logic
	var user data.User
	_ = db.Transaction(func(db data.MasterQ) error {
		user = db.Users().Insert(data.User{
			Address: ethAddress,
		})
		return nil
	})

	token, err := helpers.GenerateJWT(&user, helpers.AuthTypeSession, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a token"
		logger.WithError(err).Error(details)
		ape.RenderErr(w, errors.InternalError(details))
		return
	}
	refreshToken, err := helpers.GenerateRefreshToken(&user, helpers.ServiceConfig(r))
	if err != nil {
		details := "failed to generate a refresh token"
		logger.WithError(err).Error(details)
		ape.RenderErr(w, errors.InternalError(details))
		return
	}
	// db.Nonce().FilterByAddress(ethAddress).Delete()

	result := models.NewRegistrationModel(
		token,
		refreshToken,
		helpers.ServiceConfig(r).TokenExpireTime,
		helpers.ServiceConfig(r).RefreshTokenExpireTime,
		user,
	)

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, result)
}
