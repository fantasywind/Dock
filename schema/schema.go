package schema;

import (
  "log"
  "github.com/graphql-go/graphql"
  "dock/schema/image"
);

func Schema() *graphql.Schema {
  fields := graphql.Fields {
    "images": image.ImageListQuery,
  };

  rootQuery := graphql.ObjectConfig {
    Name: "Query",
    Fields: fields,
  };

  schemaConfig := graphql.SchemaConfig { Query: graphql.NewObject(rootQuery) }

  schema, err := graphql.NewSchema(schemaConfig);

  if err != nil {
    log.Fatalf("Failed to create new schema, error: %v", err);
  }

  return &schema;
}
