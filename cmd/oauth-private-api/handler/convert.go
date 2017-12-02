package handler

import "github.com/NeuronGroup/account-oauth/models"
import api "github.com/NeuronGroup/account-oauth/api/private/gen/models"

func fromAuthorizationCode(p *models.AuthorizationCode) (r *api.AuthorizationCode) {
	if p == nil {
		return nil
	}

	r = &api.AuthorizationCode{}
	r.Code = p.Code
	r.ExpiresSeconds = p.ExpireSeconds

	return r
}
