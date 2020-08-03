generate: go run github.com/99designs/gqlgen generate

generate data loader: go run github.com/vektah/dataloaden ReferenceLoader2 string '*graphql.model.Reference'

run project: go run server.go
