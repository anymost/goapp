package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

type query struct{
}

func main() {
	s := `
	schema {
		query: Query
	}
	type Query {
		hello: String!
		name: String!
	}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (query *query) Hello() string {
	return "Hello, world!"
}

func (query *query) Name() string  {
	return "query's name"
}
