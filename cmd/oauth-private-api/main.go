package main

import (
	"github.com/NeuronFramework/restful"
	"github.com/NeuronOauth/oauth/api-private/gen/restapi"
	"github.com/NeuronOauth/oauth/api-private/gen/restapi/operations"
	"github.com/NeuronOauth/oauth/cmd/oauth-private-api/handler"
	"github.com/go-openapi/loads"
	"net/http"
)

func main() {
	restful.Run(func() (http.Handler, error) {
		h, err := handler.NewOauthHandler()
		if err != nil {
			return nil, err
		}

		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return nil, err
		}

		api := operations.NewOauthPrivateAPI(swaggerSpec)
		api.AuthorizeHandler = operations.AuthorizeHandlerFunc(h.Authorize)

		return api.Serve(nil), nil
	})
}
