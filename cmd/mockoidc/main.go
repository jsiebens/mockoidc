package main

import (
	"github.com/labstack/echo/v4"
	"github.com/oauth2-proxy/mockoidc"
	"github.com/oauth2-proxy/mockoidc/pkg/gen/mockoidc/v1/mockoidcv1connect"
	"github.com/oauth2-proxy/mockoidc/pkg/service"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	if err := serverCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func serverCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mockoidc",
		SilenceUsage: true,
	}

	var addr string
	var clientId string
	var clientSecret string
	var serverUrl string

	cmd.Flags().StringVarP(&addr, "listen-addr", "", ":8080", "")
	cmd.Flags().StringVarP(&clientId, "client-id", "", "foo", "")
	cmd.Flags().StringVarP(&clientSecret, "client-secret", "", "bar", "")
	cmd.Flags().StringVarP(&serverUrl, "server-url", "", "http://localhost:8080", "")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		oidc, err := getMockOIDC(clientId, clientSecret, serverUrl)
		if err != nil {
			return err
		}

		oidcPath, oidcHandler := oidc.Handler()
		servicePath, serviceHandler := mockoidcv1connect.NewMockOIDCServiceHandler(service.New(oidc.UserQueue))

		e := echo.New()
		e.HideBanner = true
		e.Any(oidcPath+"/*", echo.WrapHandler(oidcHandler))
		e.Any(servicePath+"*", echo.WrapHandler(serviceHandler))
		return e.Start(addr)
	}

	return cmd
}

func getMockOIDC(clientId string, clientSecret string, serverUrl string) (*mockoidc.MockOIDC, error) {
	server, err := mockoidc.NewServer(nil)
	if err != nil {
		return nil, err
	}

	server.ServerURL = serverUrl
	server.ClientID = clientId
	server.ClientSecret = clientSecret

	return server, nil
}
