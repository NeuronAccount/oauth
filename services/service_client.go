package services

import (
	"context"
	"fmt"
	"github.com/NeuronAccount/oauth/models"
	"github.com/NeuronAccount/oauth/storages/oauth_db"
)

func (s *OauthService) ClientLogin(ctx context.Context, clientId string, password string) (c *models.OauthClient, err error) {
	dbClient, err := s.oauthDB.OauthClient.GetQuery().ClientId_Equal(clientId).QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbClient == nil {
		return nil, fmt.Errorf("clientId not exists")
	}

	if dbClient.PasswordHash != password {
		return nil, fmt.Errorf("password failed")
	}

	return oauth_db.FromOauthClient(dbClient), nil
}
