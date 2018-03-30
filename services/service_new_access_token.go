package services

import (
	"github.com/NeuronFramework/rand"
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/models"
	"github.com/NeuronOauth/oauth/storages/oauth_db"
)

func (s *OauthService) newAccessToken(ctx *restful.Context, clientId string, accountId string, scope string) (accessToken *models.AccessToken, err error) {
	dbAccessToken := &oauth_db.AccessToken{}
	dbAccessToken.AccessToken = rand.NextHex(16)
	dbAccessToken.ClientId = clientId
	dbAccessToken.AccountId = accountId
	dbAccessToken.OauthScope = scope
	dbAccessToken.ExpireSeconds = 300
	_, err = s.oauthDB.AccessToken.Insert(ctx, nil, dbAccessToken)
	if err != nil {
		return nil, err
	}

	dbRefreshToken := &oauth_db.RefreshToken{}
	dbRefreshToken.RefreshToken = rand.NextHex(16)
	dbRefreshToken.ClientId = clientId
	dbRefreshToken.AccountId = accountId
	dbRefreshToken.OauthScope = scope
	dbRefreshToken.ExpireSeconds = 300
	_, err = s.oauthDB.RefreshToken.Insert(ctx, nil, dbRefreshToken)
	if err != nil {
		return nil, err
	}

	accessToken = oauth_db.FromAccessToken(dbAccessToken)
	accessToken.RefreshToken = dbRefreshToken.RefreshToken

	accessToken.TokenType = "bearer"

	return accessToken, err
}
