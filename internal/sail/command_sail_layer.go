// Package sail provides commands to create layers for a bounded context.
package sail

const (
	// allTransportLayer defines all layers
	allTransportLayer string = "all"
	// webTransportLayer defines the web layer
	webTransportLayer string = "web"
	// restTransportLayer defines the rest layer
	restTransportLayer string = "rest"
	// gRPCTransportLayer defines the grpc layer
	gRPCTransportLayer string = "grpc"
	// GRPCLayer defines the graphql layer
	graphQLTransportLayer string = "graphql"
	// webhookTransportLayer defines the webhook layer
	webhookTransportLayer string = "webhook"
)

// TypeFuncCreateLayerTransport defines a function type for creating transport layers.
type TypeFuncCreateLayerTransport func(*SailTransport)

// mapperCreateLayerTransport maps layers to their corresponding creation functions.
var mapperCreateLayerTransport = map[string]TypeFuncCreateLayerTransport{
	webTransportLayer:     createTransport,
	restTransportLayer:    createTransport,
	gRPCTransportLayer:    createGRPCTransport,
	graphQLTransportLayer: createGraphQLTransport,
	webhookTransportLayer: createWebhookTransport,
	allTransportLayer:     createAllTransport,
}
