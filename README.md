# 🛶 Jangada

**Lightning-fast web framework for Go, inspired by Ruby on Rails**

Jangada brings the developer happiness and convention-over-configuration philosophy of Ruby on Rails to the speed and reliability of Go. Build web applications faster than ever while maintaining the performance that only Go can deliver.

## Why Jangada?

- **Rails-inspired**: Familiar MVC architecture with Go's type safety
- **Blazing Fast**: Native Go performance with sub-millisecond response times  
- **Batteries Included**: ORM, routing, templates, CLI tools - everything you need
- **Convention over Configuration**: Sensible defaults, minimal boilerplate
- **Developer Experience**: Hot reloading, powerful generators, intuitive API
- **Production Ready**: Built-in security, monitoring, and scalability features

## Philosophy

Jangada believes that web development should be both **fast to write** and **fast to run**. We combine Rails' developer ergonomics with Go's legendary performance, giving you:

- Rapid prototyping capabilities
- Enterprise-grade scalability  
- Type-safe development
- Predictable, maintainable code

## Quick Start

### Installation

```bash
go install github.com/isaqueveras/jangada/cmd/jangada@latest
```

## Features

### Core Features
- **RESTful Routing**: Automatic CRUD routes with `Resources()`
- **Active Record ORM**: Intuitive database interactions
- **MVC Architecture**: Clean separation of concerns
- **Template Engine**: Rails-like templating system
- **Middleware Support**: Pluggable request/response processing
- **JSON/XML APIs**: First-class API development support

### CLI Tools
- **Generators**: `jangada generate model User`, `jangada generate controller Posts`
- **Database Migrations**: Version-controlled schema changes
- **Server**: Built-in development server with hot reload
- **Console**: Interactive REPL for your application

### Advanced Features
- **Hot Reloading**: Automatic server restart on code changes
- **Authentication**: Built-in user authentication system
- **Security**: CSRF protection, secure headers, input validation
- **Monitoring**: Built-in metrics and health checks
- **Deployment**: One-command deployment to various platforms

## Project Structure

Jangada follows familiar Rails conventions:

```
myapp/
├── app/
│   ├── controllers/
│   ├── models/
│   ├── views/
│   └── middleware/
├── config/
│   ├── database.go
│   ├── routes.go
│   └── app.go
├── db/
│   └── migrations/
├── public/
│   ├── css/
│   ├── js/
│   └── images/
├── test/
└── main.go
```

## Configuration

```go
// config/app.go
func Configure() *jangada.Config {
    return &jangada.Config{
        Environment: jangada.GetEnv("jangada_ENV", "development"),
        Port:        jangada.GetEnv("PORT", "3000"),
        Database: &jangada.DatabaseConfig{
            Driver: "postgres",
            URL:    jangada.GetEnv("DATABASE_URL", "postgres://..."),
        },
        Cache: &jangada.CacheConfig{
            Driver: "redis",
            URL:    jangada.GetEnv("REDIS_URL", "redis://localhost:6379"),
        },
    }
}
```

## Roadmap

- [ ] **v0.1**: Core MVC, routing, basic ORM
- [ ] **v0.2**: CLI tools, generators, migrations  
- [ ] **v0.3**: Template engine, asset pipeline
- [ ] **v0.4**: Authentication, security middleware
- [ ] **v0.5**: Real-time features (WebSockets)
- [ ] **v1.0**: Production ready, comprehensive documentation

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [Ruby on Rails](https://rubyonrails.org/) - for showing us that developer happiness matters
- Built with [Go](https://golang.org/) - for proving that performance and simplicity can coexist
- Thanks to all our [contributors](https://github.com/isaqueveras/jangada/graphs/contributors)

---

**Made with 🛶 by developers who believe that web frameworks should be both fast and fun.**

[Documentation](https://jangada-framework.dev) | [Examples](https://github.com/isaqueveras/jangada-examples) 
