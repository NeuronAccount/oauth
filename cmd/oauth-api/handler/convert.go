package handler

import (
	api "github.com/NeuronAccount/oauth/api/oauth/gen/models"
	"github.com/NeuronAccount/oauth/models"
)

func fromTokenResponse(p *models.AccessToken) (r *api.AccessToken) {
	if p == nil {
		return nil
	}

	r = &api.AccessToken{}
	r.TokenType = p.TokenType
	r.AccessToken = p.AccessToken
	r.ExpiresIn = p.ExpiresIn
	r.RefreshToken = p.RefreshToken
	r.Scope = p.Scope

	return r
}
