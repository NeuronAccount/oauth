package handler

import (
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/api-private/gen/restapi/operations"
	"github.com/NeuronOauth/oauth/models"
	"github.com/NeuronOauth/oauth/services"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type OauthHandler struct {
	logger  *zap.Logger
	service *services.OauthService
}

func NewOauthHandler() (h *OauthHandler, err error) {
	h = &OauthHandler{}
	h.logger = log.TypedLogger(h)
	h.service, err = services.NewOauthService(&services.OauthServiceOptions{})
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *OauthHandler) Authorize(p operations.AuthorizeParams) middleware.Responder {
	authorizationCode, err := h.service.Authorize(restful.NewContext(p.HTTPRequest), &models.AuthorizeParams{
		AccountJwt:   p.AccountJwt,
		ClientID:     p.ClientID,
		RedirectURI:  p.RedirectURI,
		ResponseType: p.ResponseType,
		State:        p.State,
		Scope:        p.Scope,
	})

	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewAuthorizeOK().WithPayload(fromAuthorizationCode(authorizationCode))
}
