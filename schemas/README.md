# go-cloudflareanalytics / schemas

Schema packages implement the GraphQL schema as Go structures for the `cloudflareanalytics` module to make use of.

## About schema packages

Schema packages implement (at least) two types:
- `Data`: holding the GraphQL schema, as discovered through [GraphQL introspection](https://developers.cloudflare.com/analytics/graphql-api/getting-started/explore-graphql-schema);
- `Errors`: holding the error array format, as returned by the API.

Additionally, schema packages should export the canonical GraphQL API endpoint for their particular version as `Endpoint`.

Each API version (as reported in the API endpoint path) should have its own schema package.

## Schema lifecycle

Schema packages' data structures are generated through [GraphQL introspection](https://developers.cloudflare.com/analytics/graphql-api/getting-started/explore-graphql-schema), with the exceptions of `beta` and `deprecated` data sets.

As data sets mature out of `beta` status they should get introduced into their schemas.

Data sets that are deprecated _after_ their schema package has been published, must not be removed, to avoid breaking API compatibility.
Beware, though, that Cloudflare may terminate them on the server-side, breaking compatibility.

## Using multiple schemas (not recommended)

Multiple schema packages shall not be imported at the same time, therefore they are all named `schema`.

If that were a problem and multiple must be imported at the same time, they can be renamed when importing:

```go
import (
	schema_v3 "github.com/ricardbejarano/go-cloudflareanalytics/schemas/v3"
	schema_v4 "github.com/ricardbejarano/go-cloudflareanalytics/schemas/v4"
)
```
