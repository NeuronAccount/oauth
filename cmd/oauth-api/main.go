package main

import (
	"github.com/NeuronAccount/oauth/api/gen/restapi"
	"github.com/NeuronAccount/oauth/api/gen/restapi/operations"
	"github.com/NeuronAccount/oauth/cmd/oauth-api/handler"
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

	var bindAddr string

	cmd := cobra.Command{}
	cmd.PersistentFlags().StringVar(&bindAddr, "bind-addr", ":8084", "api server bind addr")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
		if err != nil {
			return err
		}
		api := operations.NewOauthAPI(swaggerSpec)

		h, err := handler.NewOauthHandler()
		if err != nil {
			return err
		}

		api.BasicAuth = h.BasicAuth

		api.TokenHandler = operations.TokenHandlerFunc(h.Token)
		api.MeHandler = operations.MeHandlerFunc(h.Me)

		logger.Info("Start server", zap.String("addr", bindAddr))
		err = http.ListenAndServe(bindAddr,
			restful.Recovery(cors.AllowAll().Handler(api.Serve(nil))))
		if err != nil {
			return err
		}

		return nil
	}
	cmd.Execute()
}
