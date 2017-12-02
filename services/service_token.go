package services

import (
	"context"
	"github.com/NeuronFramework/rand"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/storages/oauth"
)

func (s *OauthService) newAccessToken(clientId string, accountId string, scope string) (accessToken *models.AccessToken, err error) {
	dbAccessToken := &oauth.AccessToken{}
	dbAccessToken.AccessToken = rand.NextBase64(16)
	dbAccessToken.ClientId = clientId
	dbAccessToken.AccountId = accountId
	dbAccessToken.OauthScope = scope
	dbAccessToken.ExpireSeconds = 300
	_, err = s.db.AccessToken.Insert(context.Background(), nil, dbAccessToken)
	if err != nil {
		return nil, err
	}

	dbRefreshToken := &oauth.RefreshToken{}
	dbRefreshToken.RefreshToken = rand.NextBase64(16)
	dbRefreshToken.ClientId = clientId
	dbRefreshToken.AccountId = accountId
	dbRefreshToken.OauthScope = scope
	dbRefreshToken.ExpireSeconds = 300
	_, err = s.db.RefreshToken.Insert(context.Background(), nil, dbRefreshToken)
	if err != nil {
		return nil, err
	}

	accessToken = oauth.FromAccessToken(dbAccessToken)
	accessToken.RefreshToken = dbRefreshToken.RefreshToken

	accessToken.TokenType = "bearer"

	return accessToken, err
}
