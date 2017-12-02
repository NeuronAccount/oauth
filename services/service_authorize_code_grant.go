package services

import (
	"context"
	"github.com/NeuronFramework/errors"
	"github.com/NeuronAccount/oauth/models"
)

func (s *OauthService) AuthorizeCodeGrant(authorizationCode string, redirectUri string, clientId string, oAuth2Client *models.OauthClient) (accessToken *models.AccessToken, err error) {
	dbAuthorizationCode, err := s.db.AuthorizationCode.GetQuery().
		AuthorizationCode_Equal(authorizationCode).
		QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	if dbAuthorizationCode == nil {
		return nil, errors.InvalidParam("AuthorizationCode", "无效的AuthorizationCode")
	}

	return s.newAccessToken(dbAuthorizationCode.ClientId, dbAuthorizationCode.AccountId, dbAuthorizationCode.OauthScope)
}
