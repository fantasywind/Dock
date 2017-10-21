package main

import (
  "fmt"
  "strconv"
  "net/http"
  "dock/schema"
  "github.com/graphql-go/handler"
);

var PORT = 9600;

func main() {
  h := handler.New(&handler.Config {
    Schema: schema.Schema(),
    Pretty: true,
    GraphiQL: true,
  });

  http.Handle("/graphql", h);

  fmt.Printf("GraphQL Server listen on port: %v", PORT);

  http.ListenAndServe(":" + strconv.Itoa(PORT), nil);
}
