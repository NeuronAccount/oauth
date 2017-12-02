package handler

import (
	"github.com/NeuronAccount/oauth/api/oauth/gen/restapi/operations"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/services"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type OauthHandlerOptions struct {
}

type OauthHandler struct {
	logger  *zap.Logger
	options *OauthHandlerOptions
	service *services.OauthService
}

func NewOauthHandler(options *OauthHandlerOptions) (h *OauthHandler, err error) {
	h = &OauthHandler{}
	h.logger = log.TypedLogger(h)
	h.options = options
	h.service, err = services.NewOauthService(&services.OauthServiceOptions{})
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *OauthHandler) BasicAuth(clientId string, password string) (interface{}, error) {
	return h.service.ClientLogin(clientId, password)
}

func (h *OauthHandler) Token(p operations.TokenParams, oauthClient interface{}) middleware.Responder {
	if oauthClient == nil {
		return restful.Responder(errors.Unauthorized("client认证失败"))
	}

	if p.GrantType == "authorization_code" {
		if p.Code == nil {
			return restful.Responder(errors.InvalidParam("Code", "不能为空"))
		}

		if p.RedirectURI == nil {
			return restful.Responder(errors.InvalidParam("RedirectURI", "不能为空"))
		}

		if p.ClientID == nil {
			return restful.Responder(errors.InvalidParam("ClientID", "不能为空"))
		}

		result, err := h.service.AuthorizeCodeGrant(*p.Code, *p.RedirectURI, *p.ClientID, oauthClient.(*models.OauthClient))
		if err != nil {
			return restful.Responder(err)
		}

		return operations.NewTokenOK().WithPayload(fromTokenResponse(result))
	} else if p.GrantType == "refresh_token" {
		if p.RefreshToken == nil {
			return restful.Responder(errors.InvalidParam("RefreshToken", "不能为空"))
		}

		if p.Scope == nil {
			return restful.Responder(errors.InvalidParam("Scope", "不能为空"))
		}

		result, err := h.service.RefreshTokenGrant(*p.RefreshToken, *p.Scope, oauthClient.(*models.OauthClient))
		if err != nil {
			return restful.Responder(err)
		}

		return operations.NewTokenOK().WithPayload(fromTokenResponse(result))
	} else {
		return restful.Responder(errors.InvalidParam("GrantType", "未知的类型"))
	}
}

func (h *OauthHandler) Me(p operations.MeParams) middleware.Responder {
	openId, err := h.service.Me(p.AccessToken)
	if err != nil {
		return restful.Responder(err)
	}

	return operations.NewMeOK().WithPayload(openId)
}
