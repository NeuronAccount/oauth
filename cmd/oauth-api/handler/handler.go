package handler

import (
	"context"
	"fmt"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/log"
	"github.com/NeuronOauth/oauth/api/gen/restapi/operations"
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

func (h *OauthHandler) BasicAuth(clientId string, password string) (interface{}, error) {
	fmt.Println("BasicAuth", clientId, password)
	c, err := h.service.ClientLogin(context.Background(), clientId, password)
	fmt.Println("BasicAuth", c, err)
	return c, err
}

func (h *OauthHandler) Token(p operations.TokenParams, oauthClient interface{}) middleware.Responder {
	fmt.Println("token", oauthClient)

	if oauthClient == nil {
		return errors.Unauthorized("client认证失败")
	}

	if p.GrantType == "authorization_code" {
		if p.Code == nil {
			return errors.InvalidParam("Code", "不能为空")
		}

		if p.RedirectURI == nil {
			return errors.InvalidParam("RedirectURI", "不能为空")
		}

		if p.ClientID == nil {
			return errors.InvalidParam("ClientID", "不能为空")
		}

		result, err := h.service.AuthorizeCodeGrant(context.Background(),
			*p.Code, *p.RedirectURI, *p.ClientID, oauthClient.(*models.OauthClient))
		if err != nil {
			return errors.Wrap(err)
		}

		return operations.NewTokenOK().WithPayload(fromTokenResponse(result))
	} else if p.GrantType == "refresh_token" {
		if p.RefreshToken == nil {
			return errors.InvalidParam("RefreshToken", "不能为空")
		}

		if p.Scope == nil {
			return errors.InvalidParam("Scope", "不能为空")
		}

		result, err := h.service.RefreshTokenGrant(context.Background(),
			*p.RefreshToken, *p.Scope, oauthClient.(*models.OauthClient))
		if err != nil {
			return errors.Wrap(err)
		}

		return operations.NewTokenOK().WithPayload(fromTokenResponse(result))
	} else {
		return errors.InvalidParam("GrantType", "未知的类型")
	}
}

func (h *OauthHandler) Me(p operations.MeParams) middleware.Responder {
	openId, err := h.service.Me(context.Background(), p.AccessToken)
	if err != nil {
		return errors.Wrap(err)
	}

	return operations.NewMeOK().WithPayload(openId)
}
