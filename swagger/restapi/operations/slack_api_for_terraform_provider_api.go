// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/berquerant/terraform-slack/swagger/models"
	"github.com/berquerant/terraform-slack/swagger/restapi/operations/channel"
	"github.com/berquerant/terraform-slack/swagger/restapi/operations/health"
)

// NewSlackAPIForTerraformProviderAPI creates a new SlackAPIForTerraformProvider instance
func NewSlackAPIForTerraformProviderAPI(spec *loads.Document) *SlackAPIForTerraformProviderAPI {
	return &SlackAPIForTerraformProviderAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		HealthCheckAliveHandler: health.CheckAliveHandlerFunc(func(params health.CheckAliveParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation health.CheckAlive has not yet been implemented")
		}),
		ChannelCreateChannelHandler: channel.CreateChannelHandlerFunc(func(params channel.CreateChannelParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation channel.CreateChannel has not yet been implemented")
		}),
		ChannelFetchChannelHandler: channel.FetchChannelHandlerFunc(func(params channel.FetchChannelParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation channel.FetchChannel has not yet been implemented")
		}),
		ChannelFetchChannelsHandler: channel.FetchChannelsHandlerFunc(func(params channel.FetchChannelsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation channel.FetchChannels has not yet been implemented")
		}),
		ChannelUpdateChannelHandler: channel.UpdateChannelHandlerFunc(func(params channel.UpdateChannelParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation channel.UpdateChannel has not yet been implemented")
		}),

		// Applies when the "X-API-Key" header is set
		APIKeyAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (ApiKey) X-API-Key from header param [X-API-Key] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*SlackAPIForTerraformProviderAPI the slack API for terraform provider API */
type SlackAPIForTerraformProviderAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-API-Key provided in the header
	APIKeyAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// HealthCheckAliveHandler sets the operation handler for the check alive operation
	HealthCheckAliveHandler health.CheckAliveHandler
	// ChannelCreateChannelHandler sets the operation handler for the create channel operation
	ChannelCreateChannelHandler channel.CreateChannelHandler
	// ChannelFetchChannelHandler sets the operation handler for the fetch channel operation
	ChannelFetchChannelHandler channel.FetchChannelHandler
	// ChannelFetchChannelsHandler sets the operation handler for the fetch channels operation
	ChannelFetchChannelsHandler channel.FetchChannelsHandler
	// ChannelUpdateChannelHandler sets the operation handler for the update channel operation
	ChannelUpdateChannelHandler channel.UpdateChannelHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *SlackAPIForTerraformProviderAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SlackAPIForTerraformProviderAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SlackAPIForTerraformProviderAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SlackAPIForTerraformProviderAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SlackAPIForTerraformProviderAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SlackAPIForTerraformProviderAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SlackAPIForTerraformProviderAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SlackAPIForTerraformProviderAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SlackAPIForTerraformProviderAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SlackAPIForTerraformProviderAPI
func (o *SlackAPIForTerraformProviderAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "XAPIKeyAuth")
	}

	if o.HealthCheckAliveHandler == nil {
		unregistered = append(unregistered, "health.CheckAliveHandler")
	}
	if o.ChannelCreateChannelHandler == nil {
		unregistered = append(unregistered, "channel.CreateChannelHandler")
	}
	if o.ChannelFetchChannelHandler == nil {
		unregistered = append(unregistered, "channel.FetchChannelHandler")
	}
	if o.ChannelFetchChannelsHandler == nil {
		unregistered = append(unregistered, "channel.FetchChannelsHandler")
	}
	if o.ChannelUpdateChannelHandler == nil {
		unregistered = append(unregistered, "channel.UpdateChannelHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SlackAPIForTerraformProviderAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SlackAPIForTerraformProviderAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "ApiKey":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.APIKeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *SlackAPIForTerraformProviderAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SlackAPIForTerraformProviderAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SlackAPIForTerraformProviderAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SlackAPIForTerraformProviderAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the slack API for terraform provider API
func (o *SlackAPIForTerraformProviderAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SlackAPIForTerraformProviderAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = health.NewCheckAlive(o.context, o.HealthCheckAliveHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/channel"] = channel.NewCreateChannel(o.context, o.ChannelCreateChannelHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/channel"] = channel.NewFetchChannel(o.context, o.ChannelFetchChannelHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/channels"] = channel.NewFetchChannels(o.context, o.ChannelFetchChannelsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/channel"] = channel.NewUpdateChannel(o.context, o.ChannelUpdateChannelHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SlackAPIForTerraformProviderAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SlackAPIForTerraformProviderAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SlackAPIForTerraformProviderAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SlackAPIForTerraformProviderAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SlackAPIForTerraformProviderAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
