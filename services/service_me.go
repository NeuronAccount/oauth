package services

import (
	"context"
	"github.com/NeuronFramework/errors"
)

func (s *OauthService) Me(accessToken string) (openId string, err error) {
	dbAccessToken, err := s.db.AccessToken.GetQuery().AccessToken_Equal(accessToken).QueryOne(context.Background(), nil)
	if err != nil {
		return "", err
	}

	if dbAccessToken == nil {
		return "", errors.NotFound("accessToken不存在")
	}

	return dbAccessToken.AccountId, nil
}
