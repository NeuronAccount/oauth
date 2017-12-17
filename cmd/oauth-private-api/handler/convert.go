package handler

import "github.com/NeuronOauth/oauth/models"
import api "github.com/NeuronOauth/oauth/api-private/gen/models"

func fromAuthorizationCode(p *models.AuthorizationCode) (r *api.AuthorizationCode) {
	if p == nil {
		return nil
	}

	r = &api.AuthorizationCode{}
	r.Code = p.Code
	r.ExpiresSeconds = p.ExpireSeconds

	return r
}
