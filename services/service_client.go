package services

import (
	"fmt"
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/models"
	"github.com/NeuronOauth/oauth/storages/oauth_db"
)

func (s *OauthService) ClientLogin(ctx *restful.Context, clientId string, password string) (c *models.OauthClient, err error) {
	dbClient, err := s.oauthDB.OauthClient.GetQuery().ClientId_Equal(clientId).QueryOne(ctx, nil)
	if err != nil {
		return nil, err
	}

	if dbClient == nil {
		return nil, fmt.Errorf("clientId不存在")
	}

	if dbClient.PasswordHash != password {
		return nil, fmt.Errorf("password错误")
	}

	return oauth_db.FromOauthClient(dbClient), nil
}
