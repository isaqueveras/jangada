// Package sail provides commands to create layers for a bounded context.
package sail

// InterfaceLayer defines the type for interface layers.
type InterfaceLayer string

const (
	// AllInterfaceLayer defines all layers
	AllInterfaceLayer InterfaceLayer = "all"
	// WebInterfaceLayer defines the web layer
	WebInterfaceLayer InterfaceLayer = "web"
	// RestInterfaceLayer defines the http layer
	RestInterfaceLayer InterfaceLayer = "http"
	// GRPCInterfaceLayer defines the grpc layer
	GRPCInterfaceLayer InterfaceLayer = "grpc"
	// GRPCLayer defines the graphql layer
	GraphQLInterfaceLayer InterfaceLayer = "graphql"
	// WebhookInterfaceLayer defines the webhook layer
	WebhookInterfaceLayer InterfaceLayer = "webhook"
)

// TypeFuncCreateLayerInterface defines a function type for creating interface layers.
type TypeFuncCreateLayerInterface func(*SailInterface)

// mapperCreateLayerInterface maps layers to their corresponding creation functions.
var mapperCreateLayerInterface = map[InterfaceLayer]TypeFuncCreateLayerInterface{
	WebInterfaceLayer:     createWebInterface,
	RestInterfaceLayer:    createRestInterface,
	GRPCInterfaceLayer:    createGRPCInterface,
	GraphQLInterfaceLayer: createGraphQLInterface,
	WebhookInterfaceLayer: createWebhookInterface,
	AllInterfaceLayer:     createAllInterface,
}
