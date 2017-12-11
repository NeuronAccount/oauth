package services

import (
	"context"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronFramework/errors"
)

func (s *OauthService) AuthorizeCodeGrant(ctx context.Context, authorizationCode string, redirectUri string, clientId string, oAuth2Client *models.OauthClient) (accessToken *models.AccessToken, err error) {
	dbAuthorizationCode, err := s.oauthDB.AuthorizationCode.GetQuery().
		AuthorizationCode_Equal(authorizationCode).
		QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbAuthorizationCode == nil {
		return nil, errors.InvalidParam("AuthorizationCode", "无效的AuthorizationCode")
	}

	return s.newAccessToken(ctx, dbAuthorizationCode.ClientId, dbAuthorizationCode.AccountId, dbAuthorizationCode.OauthScope)
}
