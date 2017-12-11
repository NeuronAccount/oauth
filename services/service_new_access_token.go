package services

import (
	"context"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/storages/oauth_db"
	"github.com/NeuronFramework/rand"
)

func (s *OauthService) newAccessToken(ctx context.Context, clientId string, accountId string, scope string) (accessToken *models.AccessToken, err error) {
	dbAccessToken := &oauth_db.AccessToken{}
	dbAccessToken.AccessToken = rand.NextBase64(16)
	dbAccessToken.ClientId = clientId
	dbAccessToken.AccountId = accountId
	dbAccessToken.OauthScope = scope
	dbAccessToken.ExpireSeconds = 300
	_, err = s.oauthDB.AccessToken.Insert(ctx, nil, dbAccessToken)
	if err != nil {
		return nil, err
	}

	dbRefreshToken := &oauth_db.RefreshToken{}
	dbRefreshToken.RefreshToken = rand.NextBase64(16)
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
