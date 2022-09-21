package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/nonce-auth-svc/internal/service/helpers"
	"gitlab.com/tokene/nonce-auth-svc/resources"
)

func CreatedAt(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	db := helpers.DB(r)
	doorman := helpers.DoormanConnector(r)

	address, _, _, err := helpers.ValidatePurposeJWT(doorman, r)
	if err != nil {
		logger.WithError(err).Debug("failed authentication")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	user, err := db.Users().FilterByAddress(address).Get()
	if user == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// success logic

	result := resources.CreatedAt{
		Address:   address,
		CreatedAt: user.CreatedAt.Format("02-01-2006 15:04:05"),
	}

	ape.Render(w, result)
}
