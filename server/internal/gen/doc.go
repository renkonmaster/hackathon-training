package gen

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.5.0 -generate types,server -package gen -o api.gen.go ../../../docs/openapi.yaml
