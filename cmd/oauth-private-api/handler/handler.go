package handler

import (
	"context"
	"github.com/NeuronAccount/oauth/api-private/gen/restapi/operations"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/services"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
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
	authorizationCode, err := h.service.Authorize(context.Background(), &models.AuthorizeParams{
		AccountJwt:   p.AccountJwt,
		ClientID:     p.ClientID,
		RedirectURI:  p.RedirectURI,
		ResponseType: p.ResponseType,
		State:        p.State,
		Scope:        p.Scope,
	})

	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewAuthorizeOK().WithPayload(fromAuthorizationCode(authorizationCode))
}
