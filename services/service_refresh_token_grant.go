package services

import (
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/models"
)

func (s *OauthService) RefreshTokenGrant(ctx *restful.Context, refreshToken string, scope string, client *models.OauthClient) (accessToken *models.AccessToken, err error) {
	dbRefreshToken, err := s.oauthDB.RefreshToken.GetQuery().RefreshToken_Equal(refreshToken).QueryOne(ctx, nil)
	if err != nil {
		return nil, nil
	}

	if dbRefreshToken == nil {
		return nil, errors.InvalidParam("无效的RefreshToken")
	}

	return s.newAccessToken(ctx, dbRefreshToken.ClientId, dbRefreshToken.AccountId, dbRefreshToken.OauthScope)
}
