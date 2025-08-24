[![Go Reference](https://pkg.go.dev/badge/github.com/isaqueveras/jangada.svg)](https://pkg.go.dev/github.com/isaqueveras/jangada)
[![Build](https://github.com/isaqueveras/jangada/actions/workflows/build.yml/badge.svg)](https://github.com/isaqueveras/jangada/actions/workflows/build.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

**🛶 Jangada** is a **Go** web framework full stack, designed for productivity and **DDD best practices**. 
> Build web apps and APIs quickly with plugins, background jobs, and workflows.

---

## Features

- DDD-based structure: `application`, `domain`, `interface`, `infrastructure` and `tests`
- Plugin system via CLI
- Background jobs and chained workflows
- Integrations with Postgres, Gin, Synk, and more
- Simple YAML configuration
- Middleware and helpers for logging and request tracing

---

## Quick Start

```bash
# Install Jangada
go install github.com/isaqueveras/jangada@latest
```

## Create a new app
```bash
jangada new myapp
cd myapp
jangada serve
```
