package main

import "fmt";
import "strconv";
import "net/http";
import "dock/schema";
import "github.com/graphql-go/handler";

var PORT = 9600;

func main() {
  h := handler.New(&handler.Config {
    Schema: schema.Schema(),
    Pretty: true,
    GraphiQL: true,
  });

  http.Handle("/graphql", h);
  http.ListenAndServe(":" + strconv.Itoa(PORT), nil);

  fmt.Println("Server Listen on port: %d", PORT);
}
