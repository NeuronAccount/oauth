// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// AuthorizeURL generates an URL for the authorize operation
type AuthorizeURL struct {
	AccountJwt   string
	ClientID     string
	RedirectURI  string
	ResponseType string
	Scope        string
	State        string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AuthorizeURL) WithBasePath(bp string) *AuthorizeURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AuthorizeURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AuthorizeURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/authorize"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api-private/v1/oauth"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	accountJwt := o.AccountJwt
	if accountJwt != "" {
		qs.Set("accountJwt", accountJwt)
	}

	clientID := o.ClientID
	if clientID != "" {
		qs.Set("client_id", clientID)
	}

	redirectURI := o.RedirectURI
	if redirectURI != "" {
		qs.Set("redirect_uri", redirectURI)
	}

	responseType := o.ResponseType
	if responseType != "" {
		qs.Set("response_type", responseType)
	}

	scope := o.Scope
	if scope != "" {
		qs.Set("scope", scope)
	}

	state := o.State
	if state != "" {
		qs.Set("state", state)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AuthorizeURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AuthorizeURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AuthorizeURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AuthorizeURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AuthorizeURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *AuthorizeURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
