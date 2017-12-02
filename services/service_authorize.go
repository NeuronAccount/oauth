package services

import (
	"context"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/storages/oauth"
	"github.com/NeuronFramework/rand"
)

func (s *OauthService) Authorize(p *models.AuthorizeParams) (code *models.AuthorizationCode, err error) {
	dbAuthorizationCode := &oauth.AuthorizationCode{}
	dbAuthorizationCode.AuthorizationCode = rand.NextBase64(16)
	dbAuthorizationCode.ClientId = p.ClientID
	dbAuthorizationCode.AccountId = p.Jwt
	dbAuthorizationCode.RedirectUri = p.RedirectURI
	dbAuthorizationCode.OauthScope = p.Scope
	dbAuthorizationCode.ExpireSeconds = 300
	_, err = s.db.AuthorizationCode.Insert(context.Background(), nil, dbAuthorizationCode)
	if err != nil {
		return nil, err
	}

	return oauth.FromAuthorizationCode(dbAuthorizationCode), nil
}
