package handler

import "github.com/NeuronAccount/oauth/models"
import api "github.com/NeuronAccount/oauth/api-private/gen/models"

func fromAuthorizationCode(p *models.AuthorizationCode) (r *api.AuthorizationCode) {
	if p == nil {
		return nil
	}

	r = &api.AuthorizationCode{}
	r.Code = p.Code
	r.ExpiresSeconds = p.ExpireSeconds

	return r
}
