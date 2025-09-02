// Package sail provides commands to create layers for a bounded context.
package sail

// exampleCreateInterfaceText defines the example text for the 'sail interface' command.
const exampleCreateInterfaceText = `
# Create web interface layer structure for the "user" bounded context
jangada sail interface user web

# Create rest interface layer structure for the "user" bounded context
jangada sail interface user rest

# Create gRPC interface layer structure for the "user" bounded context
jangada sail interface user grpc`
