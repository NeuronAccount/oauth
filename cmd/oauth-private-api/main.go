package main

import (
	"github.com/NeuronAccount/oauth/api/private/gen/restapi"
	"github.com/NeuronAccount/oauth/api/private/gen/restapi/operations"
	"github.com/NeuronAccount/oauth/cmd/oauth-private-api/handler"
	"github.com/NeuronFramework/log"
	"github.com/NeuronFramework/restful"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	log.Init(true)

	middleware.Debug = false

	logger := zap.L().Named("main")

	var bind_addr string

	cmd := cobra.Command{}
	cmd.PersistentFlags().StringVar(&bind_addr, "bind-addr", ":8085", "api server bind addr")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return err
		}
		api := operations.NewOauthPrivateAPI(swaggerSpec)

		h, err := handler.NewOauthHandler(&handler.OauthHandlerOptions{})
		if err != nil {
			return err
		}

		api.AuthorizeHandler = operations.AuthorizeHandlerFunc(h.Authorize)

		logger.Info("Start server", zap.String("addr", bind_addr))
		err = http.ListenAndServe(bind_addr,
			restful.Recovery(cors.AllowAll().Handler(api.Serve(nil))))
		if err != nil {
			return err
		}

		return nil
	}
	cmd.Execute()
}
