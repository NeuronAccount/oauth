package services

import (
	"context"
	"fmt"
	"github.com/NeuronGroup/account-oauth/models"
	"github.com/NeuronGroup/account-oauth/storages/oauth"
)

func (s *OauthService) ClientLogin(clientId string, password string) (c *models.OauthClient, err error) {
	dbClient, err := s.db.OauthClient.GetQuery().ClientId_Equal(clientId).QueryOne(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	if dbClient == nil {
		return nil, fmt.Errorf("clientId not exists")
	}

	if dbClient.PasswordHash != password {
		return nil, fmt.Errorf("password failed")
	}

	return oauth.FromOauthClient(dbClient), nil
}
