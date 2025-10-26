// Package sail provides commands to create layers for a bounded context.
package sail

const (
	// AllTransportLayer defines all layers
	AllTransportLayer string = "all"
	// WebTransportLayer defines the web layer
	WebTransportLayer string = "web"
	// RestTransportLayer defines the rest layer
	RestTransportLayer string = "rest"
	// GRPCTransportLayer defines the grpc layer
	GRPCTransportLayer string = "grpc"
	// GRPCLayer defines the graphql layer
	GraphQLTransportLayer string = "graphql"
	// WebhookTransportLayer defines the webhook layer
	WebhookTransportLayer string = "webhook"
)

// TypeFuncCreateLayerTransport defines a function type for creating transport layers.
type TypeFuncCreateLayerTransport func(*SailTransport)

// mapperCreateLayerTransport maps layers to their corresponding creation functions.
var mapperCreateLayerTransport = map[string]TypeFuncCreateLayerTransport{
	WebTransportLayer:     createTransport,
	RestTransportLayer:    createTransport,
	GRPCTransportLayer:    createGRPCTransport,
	GraphQLTransportLayer: createGraphQLTransport,
	WebhookTransportLayer: createWebhookTransport,
	AllTransportLayer:     createAllTransport,
}
