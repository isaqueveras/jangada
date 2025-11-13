[![Go Reference](https://pkg.go.dev/badge/github.com/isaqueveras/jangada.svg)](https://pkg.go.dev/github.com/isaqueveras/jangada)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/isaqueveras/jangada)](https://goreportcard.com/report/github.com/isaqueveras/jangada)

**🛶 Jangada** is a **Go** web framework full stack, designed for productivity and **DDD best practices**. 
> Build web apps and APIs quickly with plugins, background jobs, and workflows.

<img width="2048" height="1536" alt="Welcome page" src="https://github.com/user-attachments/assets/d1a53384-e300-4296-a760-abc75b26e208" />


---

## Features

- DDD-based structure: `application`, `domain`, `transport`, `infrastructure` and `tests`
- Background jobs and chained workflows

---

## Install Jangada CLI

```bash
# Install Jangada
go install github.com/isaqueveras/jangada@latest
```

## Create a new app
Create the foundation of an application to build a monolithic or microservice

```bash
$ jangada new commerce

# or configure the base with the flags
$ jangada new commerce --mod=github.com/isaqueveras/commerce --host=localhost:8782 --db=postgres
```

## Create a domain layer
The domain layer is where business rules, entities, repositories, and services should be defined

```bash
$ jangada sail domain crm/customer
```

## Create a application layer
The application layer is where the orchestrators and services that create flows within the system are concentrated

```bash
$ jangada sail application crm/customer

# or create a new service in the application
$ jangada sail application crm/customer --service=GetAllCustomerByName
```

## Create a transport layer
The transport layer is where you control the rest/grpc/web routes that access the application layer

```bash
$ jangada sail transport crm/customer

# or create a new method in the controller
$ jangada sail transport crm/customer --layer=rest --name=GetAllCustomerByName
```

## Create a infrastructure layer
The infrastructure layer is where you implement the repository and external calls, such as database access or HTTP/GRPC requests

```bash
$ jangada sail infrastructure crm/customer
```
