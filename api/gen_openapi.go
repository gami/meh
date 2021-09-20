package api

//go:generate oapi-codegen -generate "chi-server" -o openapi/server.go openapi.yaml
//go:generate oapi-codegen -generate "types" -o openapi/type.go openapi.yaml
//go:generate oapi-codegen -generate "spec" -o openapi/spec.go openapi.yaml
