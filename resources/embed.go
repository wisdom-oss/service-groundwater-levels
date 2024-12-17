package resources

import "embed"

//go:embed *.sql
var QueryFiles embed.FS

//go:embed schema.graphql
var GraphQLSchema string
