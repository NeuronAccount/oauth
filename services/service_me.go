package services

import (
	"context"
	"github.com/NeuronFramework/errors"
)

func (s *OauthService) Me(ctx context.Context, accessToken string) (openId string, err error) {
	dbAccessToken, err := s.oauthDB.AccessToken.GetQuery().AccessToken_Equal(accessToken).QueryOne(ctx, nil)
	if err != nil {
		return "", err
	}

	if dbAccessToken == nil {
		return "", errors.NotFound("accessToken不存在")
	}

	return dbAccessToken.AccountId, nil
}
