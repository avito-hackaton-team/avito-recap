package http

import "github.com/avito-hackaton-team/avito-recap/backend/recap/generated/recapapi"

type API struct {
	server recapapi.Handler
}

func New(recapServer recapapi.Handler) *API {
	return &API{
		server: recapServer,
	}
}
