package services

import (
	"context"
	"github.com/NeuronFramework/rand"
	"github.com/NeuronOauth/oauth/models"
	"github.com/NeuronOauth/oauth/storages/oauth_db"
	"github.com/dgrijalva/jwt-go"
)

func (s *OauthService) Authorize(ctx context.Context, p *models.AuthorizeParams) (code *models.AuthorizationCode, err error) {
	claims := jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(p.AccountJwt, &claims, func(t *jwt.Token) (interface{}, error) { return []byte("0123456789"), nil })
	if err != nil {
		return nil, err
	}

	dbAuthorizationCode := &oauth_db.AuthorizationCode{}
	dbAuthorizationCode.AuthorizationCode = rand.NextHex(16)
	dbAuthorizationCode.ClientId = p.ClientID
	dbAuthorizationCode.AccountId = claims.Subject
	dbAuthorizationCode.RedirectUri = p.RedirectURI
	dbAuthorizationCode.OauthScope = p.Scope
	dbAuthorizationCode.ExpireSeconds = 300
	_, err = s.oauthDB.AuthorizationCode.Insert(ctx, nil, dbAuthorizationCode)
	if err != nil {
		return nil, err
	}

	return oauth_db.FromAuthorizationCode(dbAuthorizationCode), nil
}
