package oauth

import "github.com/NeuronGroup/account-oauth/models"

func FromOauthClient(p *OauthClient) (r *models.OauthClient) {
	if p == nil {
		return nil
	}

	r = &models.OauthClient{}
	r.ClientId = p.ClientId
	r.PasswordHash = p.PasswordHash
	r.AccountId = p.AccountId
	r.RedirectUri = p.RedirectUri

	return r
}

func FromAuthorizationCode(p *AuthorizationCode) (r *models.AuthorizationCode) {
	if p == nil {
		return nil
	}

	r = &models.AuthorizationCode{}
	r.Code = p.AuthorizationCode
	r.ClientId = p.ClientId
	r.AccountId = p.AccountId
	r.Scope = p.OauthScope
	r.ExpireSeconds = p.ExpireSeconds
	r.RedirectUri = p.RedirectUri

	return r
}

func FromAccessToken(p *AccessToken) (r *models.AccessToken) {
	if p == nil {
		return nil
	}

	r = &models.AccessToken{}
	r.AccessToken = p.AccessToken
	r.ClientId = p.ClientId
	r.AccountId = p.AccountId
	r.Scope = p.OauthScope
	r.ExpiresIn = p.ExpireSeconds

	return r
}
