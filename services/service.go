package services

import (
	"github.com/NeuronFramework/log"
	"github.com/NeuronOauth/oauth/storages/oauth_db"
	"go.uber.org/zap"
)

type OauthServiceOptions struct {
}

type OauthService struct {
	logger  *zap.Logger
	options *OauthServiceOptions
	oauthDB *oauth_db.DB
}

func NewOauthService(options *OauthServiceOptions) (s *OauthService, err error) {
	s = &OauthService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	s.oauthDB, err = oauth_db.NewDB()
	if err != nil {
		return nil, err
	}

	return s, nil
}
