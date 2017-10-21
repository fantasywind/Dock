package image;

import (
  "strings"
  "log"
  "github.com/graphql-go/graphql"
  "github.com/fsouza/go-dockerclient"
  "dock/schema/config"
);

var ImageRepositoryType = graphql.NewObject(graphql.ObjectConfig {
  Name: "ImageRepository",
  Fields: graphql.Fields {
    "id": &graphql.Field {
      Type: graphql.NewNonNull(graphql.String),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return p.Source.(string), nil;
      },
    },
    "name": &graphql.Field {
      Type: graphql.NewNonNull(graphql.String),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return strings.Split(p.Source.(string), ":")[0], nil;
      },
    },
    "tag": &graphql.Field {
      Type: graphql.NewNonNull(graphql.String),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return strings.Split(p.Source.(string), ":")[1], nil;
      },
    },
  },
});

var ImageType = graphql.NewObject(graphql.ObjectConfig {
  Name: "Image",
  Fields: graphql.Fields {
    "id": &graphql.Field {
      Type: graphql.NewNonNull(graphql.String),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return p.Source.(docker.APIImages).ID[7:], nil;
      },
    },
    "repositories": &graphql.Field {
      Type: graphql.NewList(ImageRepositoryType),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return p.Source.(docker.APIImages).RepoTags, nil;
      },
    },
    "createdOn": &graphql.Field {
      Type: graphql.NewNonNull(graphql.Int),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return p.Source.(docker.APIImages).Created, nil;
      },
    },
    "size": &graphql.Field {
      Type: graphql.NewNonNull(graphql.Int),
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return p.Source.(docker.APIImages).Size, nil;
      },
    },
  },
});

var ImageListQuery = &graphql.Field {
  Type: graphql.NewList(ImageType),
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    client, err := docker.NewClient(config.Endpoint);

    if err != nil {
      log.Fatalf("Failed to create connection to docker engine on %v, error: %v", config.Endpoint, err);
    }

    images, err := client.ListImages(docker.ListImagesOptions { All: false });

    if err != nil {
      log.Fatalf("Failed to list docker images error: %v", err);
    }

    return images, nil;
  },
};
