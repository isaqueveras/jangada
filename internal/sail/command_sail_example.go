// Package sail provides commands to create layers for a bounded context.
package sail

// exampleCreateTransportText defines the example text for the 'sail transport' command.
const exampleCreateTransportText = `
# Create web transport layer structure for the "user" bounded context
jangada sail transport user/user web

# Create rest transport layer structure for the "user" bounded context
jangada sail transport user/user rest

# Create gRPC transport layer structure for the "user" bounded context
jangada sail transport user/user grpc`
