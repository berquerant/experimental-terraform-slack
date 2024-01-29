// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/berquerant/terraform-slack/swagger/models"
	"github.com/berquerant/terraform-slack/swagger/restapi/operations"
	"github.com/berquerant/terraform-slack/swagger/restapi/operations/channel"
	"github.com/berquerant/terraform-slack/swagger/restapi/operations/health"
)

//go:generate ../../tools/swagger.sh do generate server --target ../../swagger --name SlackAPIForTerraformProvider --spec ../../swagger.yml --principal models.Principal

func configureFlags(api *operations.SlackAPIForTerraformProviderAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SlackAPIForTerraformProviderAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-Key" header is set
	api.APIKeyAuth = func(token string) (*models.Principal, error) {
		if token == "" {
			return nil, errors.Unauthenticated("token")
		}
		p := models.Principal(token)
		return &p, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	api.HealthCheckAliveHandler = health.CheckAliveHandlerFunc(health.HandleCheckAlive)
	api.ChannelUpdateChannelHandler = channel.UpdateChannelHandlerFunc(channel.HandleUpdateChannel)
	api.ChannelCreateChannelHandler = channel.CreateChannelHandlerFunc(channel.HandleCreateChannel)
	api.ChannelFetchChannelHandler = channel.FetchChannelHandlerFunc(channel.HandleFetchChannel)
	api.ChannelFetchChannelsHandler = channel.FetchChannelsHandlerFunc(channel.HandleFetchChannels)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
