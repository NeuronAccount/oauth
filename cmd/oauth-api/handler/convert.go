package handler

import (
	api "github.com/NeuronOauth/oauth/api/gen/models"
	"github.com/NeuronOauth/oauth/models"
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
