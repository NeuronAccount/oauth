package services

import (
	"context"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/storages/oauth_db"
	"github.com/NeuronFramework/rand"
)

func (s *OauthService) Authorize(ctx context.Context, p *models.AuthorizeParams) (code *models.AuthorizationCode, err error) {
	dbAuthorizationCode := &oauth_db.AuthorizationCode{}
	dbAuthorizationCode.AuthorizationCode = rand.NextBase64(16)
	dbAuthorizationCode.ClientId = p.ClientID
	dbAuthorizationCode.AccountId = p.Jwt
	dbAuthorizationCode.RedirectUri = p.RedirectURI
	dbAuthorizationCode.OauthScope = p.Scope
	dbAuthorizationCode.ExpireSeconds = 300
	_, err = s.oauthDB.AuthorizationCode.Insert(ctx, nil, dbAuthorizationCode)
	if err != nil {
		return nil, err
	}

	return oauth_db.FromAuthorizationCode(dbAuthorizationCode), nil
}
