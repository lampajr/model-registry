# Model Registry
Model registry provides a central repository for model developers to store and manage models, versions, and artifacts metadata. A Go-based application that leverages [ml_metadata](https://github.com/google/ml-metadata/) project under the hood.

## Red Hat's Pledge
- Red Hat drives the project's development through Open Source principles, ensuring transparency, sustainability, and community ownership.
- Red Hat values the Kubeflow community and commits to providing a minimum of 12 months' notice before ending project maintenance after the initial release.

![build checks status](https://github.com/opendatahub-io/model-registry/actions/workflows/build.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/github/opendatahub-io/model-registry/graph/badge.svg?token=61URLQA3VS)](https://codecov.io/github/opendatahub-io/model-registry)

## Pre-requisites:
- go >= 1.19
- protoc v24.3 - [Protocol Buffers v24.3 Release](https://github.com/protocolbuffers/protobuf/releases/tag/v24.3)
- npm >= 10.2.0 - [Installing Node.js and npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- Java >= 11.0
- python 3.9

## OpenAPI Proxy Server

The model registry proxy server implementation follows a contract-first approach, where the contract is identified by [model-registry.yaml](api/openapi/model-registry.yaml) OpenAPI specification.

You can also easily display the latest OpenAPI contract for model-registry in a Swagger-like editor directly from this repository; for example, [here](https://editor.swagger.io/?url=https://raw.githubusercontent.com/opendatahub-io/model-registry/main/api/openapi/model-registry.yaml).
### Starting the OpenAPI Proxy Server
Run the following command to start the OpenAPI proxy server from source:

```shell
make run/proxy
```
The proxy service implements the OpenAPI defined in [model-registry.yaml](api/openapi/model-registry.yaml) to create an Open Data Hub specific REST API on top of the existing ml-metadata server.

> **NOTE** The ml-metadata server must be running and accessible from the environment where model-registry starts up.

### Model registry logical model

For a high-level documentation of the Model Registry _logical model_, please check [this guide](./docs/logical_model.md).

## Model Registry Core

The model registry core is the layer which implements the core/business logic by interacting with the underlying ml-metadata server.
It provides a model registry domain-specific [api](internal/core/api.go) that is in charge to proxy all, appropriately transformed, requests to ml-metadata using gRPC calls.

### Model registry library

For more background on Model Registry Go core library and instructions on using it, please check [getting started guide](./docs/mr_go_library.md).

## Development

### Building
Run the following command to build the server binary:

```shell
make build
```

The generated binary uses `spf13` cmdline args. More information on using the server can be obtained by running the command:

```shell
./model-registry --help
```

Run the following command to clean the server binary, generated models and etc.:

```shell
make clean
```

### Testing

Run the following command to trigger all tests:

```shell
make test
```

or, to see the statement coverage:

```shell
make test-cover
```

### Docker Image
#### Building the docker image
The following command builds a docker image for the server with the tag `model-registry`:

```shell
docker build -t model-registry .
```

Note that the first build will be longer as it downloads the build tool dependencies.
Subsequent builds will re-use the cached tools layer.

#### Running the proxy server

> **NOTE:** ml-metadata server must be running and accessible, see more info on how to start the gRPC server in the official ml-metadata [documentation](https://github.com/google/ml-metadata).

The following command starts the proxy server:

```shell
docker run -d -p <hostname>:<port>:8080 --user <uid>:<gid> --name server model-registry proxy -n 0.0.0.0
```

Where, `<uid>`, `<gid>`, and `<host-path>` are the same as in the migrate command above.
And `<hostname>` and `<port>` are the local ip and port to use to expose the container's default `8080` listening port.
The server listens on `localhost` by default, hence the `-n 0.0.0.0` option allows the server port to be exposed.

#### Running model registry & ml-metadata

> **NOTE:** Docker compose must be installed in your environment.

There are two `docker-compose` files that make the startup of both model registry and ml-metadara easier, by simply running:

```shell
docker compose -f docker-compose[-local].yaml up
```

The main difference between the two docker compose files is that `-local` one build the model registry from source, the other one, instead, download the `latest` pushed [quay.io](https://quay.io/repository/opendatahub/model-registry?tab=tags) image.

When shutting down the docker compose, you might want to clean-up the SQLite db file generated by ML Metadata, for example `./test/config/ml-metadata/metadata.sqlite.db`

### Testing architecture

The following diagram illustrate testing strategy for the several components in Model Registry project:

![](/docs/Model%20Registry%20Testing%20areas.png)

Go layers components are tested with Unit Tests written in Go, as well as Integration Tests leveraging Testcontainers.
This allows to verify the expected "Core layer" of logical data mapping developed implemented in Go, matches technical expectations.

Python client is also tested with Unit Tests and Integration Tests written in Python.

End-to-end testing is developed with Robot Framework; this higher-lever layer of testing is used to:
- demonstrate *User Stories* from high level perspective
- demonstrate coherent logical data mapping by performing the same high level capabilities, using REST API flow Vs Python client flow,
directly checking the end results in the backend gRPC MLMD server.
