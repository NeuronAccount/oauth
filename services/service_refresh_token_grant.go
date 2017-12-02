package services

import (
	"context"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronFramework/errors"
)

func (s *OauthService) RefreshTokenGrant(refresh_token string, scope string, client *models.OauthClient) (accessToken *models.AccessToken, err error) {
	dbRefreshToken, err := s.db.RefreshToken.GetQuery().RefreshToken_Equal(refresh_token).QueryOne(context.Background(), nil)
	if err != nil {
		return nil, nil
	}

	if dbRefreshToken == nil {
		return nil, errors.InvalidParam("refresh_token", "无效的RefreshToken")
	}

	return s.newAccessToken(dbRefreshToken.ClientId, dbRefreshToken.AccountId, dbRefreshToken.OauthScope)
}
