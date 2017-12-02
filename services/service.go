package services

import (
	"github.com/NeuronFramework/log"
	"github.com/NeuronAccount/oauth/storages/oauth"
	"go.uber.org/zap"
)

type OauthServiceOptions struct {
}

type OauthService struct {
	logger  *zap.Logger
	options *OauthServiceOptions
	db      *oauth.DB
}

func NewOauthService(options *OauthServiceOptions) (s *OauthService, err error) {
	s = &OauthService{}
	s.logger = log.TypedLogger(s)
	s.options = options
	s.db, err = oauth.NewDB("root:123456@tcp(127.0.0.1:3307)/account-oauth?parseTime=true")
	if err != nil {
		return nil, err
	}

	return s, nil
}
