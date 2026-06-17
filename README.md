# Petstore Extended API

This repository contains a complete API generated using ogen from the petstore-extended OpenAPI v3 spec file. The code generation creates only the API without the client CLI.

## What has been done so far:
1. Downloaded the `petstore-expanded.yml` spec file into `internal/api`.
2. Initialized a Go module for the project (`petstore`).
3. Installed `mise` for tool setup and task management.
4. Created a `mise` configuration file in `.config/mise/config.toml` that sets up Go version `1.23.0` and defines `build`, `generate`, and `test` tasks.
5. Created an `ogen.yml` file to disable client generation (`paths/client`, `webhooks/client`).
6. Generated the initial boilerplate for the API using ogen in the `internal/api` directory (excluding the client generation as per requirements).
7. Ran `go mod tidy` and `go build ./...` to manage dependencies and verify build.

## What has to be done next:
1. Implement the API server to serve the generated handlers by fulfilling the handler interface.
2. Implement the API handlers with dummy data or database logic (for `findPets`, `addPet`, `find pet by id`, `deletePet`).
3. Create a main application to run the server.
4. Write tests for the implementation.
