// Package sail provides commands to create layers for a bounded context.
package sail

// TransportLayer defines the type for transport layers.
type TransportLayer string

const (
	// AllTransportLayer defines all layers
	AllTransportLayer TransportLayer = "all"
	// WebTransportLayer defines the web layer
	WebTransportLayer TransportLayer = "web"
	// RestTransportLayer defines the rest layer
	RestTransportLayer TransportLayer = "rest"
	// GRPCTransportLayer defines the grpc layer
	GRPCTransportLayer TransportLayer = "grpc"
	// GRPCLayer defines the graphql layer
	GraphQLTransportLayer TransportLayer = "graphql"
	// WebhookTransportLayer defines the webhook layer
	WebhookTransportLayer TransportLayer = "webhook"
)

// TypeFuncCreateLayerTransport defines a function type for creating transport layers.
type TypeFuncCreateLayerTransport func(*SailTransport)

// mapperCreateLayerTransport maps layers to their corresponding creation functions.
var mapperCreateLayerTransport = map[TransportLayer]TypeFuncCreateLayerTransport{
	WebTransportLayer:     createTransport,
	RestTransportLayer:    createTransport,
	GRPCTransportLayer:    createGRPCTransport,
	GraphQLTransportLayer: createGraphQLTransport,
	WebhookTransportLayer: createWebhookTransport,
	AllTransportLayer:     createAllTransport,
}
