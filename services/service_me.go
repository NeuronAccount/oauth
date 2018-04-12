package services

import (
	"github.com/NeuronFramework/errors"
	"github.com/NeuronFramework/restful"
)

func (s *OauthService) Me(ctx *restful.Context, accessToken string) (accountId string, err error) {
	dbAccessToken, err := s.oauthDB.AccessToken.GetQuery().AccessToken_Equal(accessToken).QueryOne(ctx, nil)
	if err != nil {
		return "", err
	}

	if dbAccessToken == nil {
		return "", errors.NotFound("accessToken不存在")
	}

	return dbAccessToken.AccountId, nil
}
