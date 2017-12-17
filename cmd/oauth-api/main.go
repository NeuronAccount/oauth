package main

import (
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/api/gen/restapi"
	"github.com/NeuronOauth/oauth/api/gen/restapi/operations"
	"github.com/NeuronOauth/oauth/cmd/oauth-api/handler"
	"github.com/go-openapi/loads"
	"net/http"
	"os"
)

func main() {
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "8084")

	restful.Run(func() (http.Handler, error) {
		h, err := handler.NewOauthHandler()
		if err != nil {
			return nil, err
		}

		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return nil, err
		}

		api := operations.NewOauthAPI(swaggerSpec)
		api.BasicAuth = h.BasicAuth
		api.TokenHandler = operations.TokenHandlerFunc(h.Token)
		api.MeHandler = operations.MeHandlerFunc(h.Me)

		return api.Serve(nil), nil
	})
}
