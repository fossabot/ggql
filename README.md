# GGql

A GraphQL implementation for a GraphQL API that is easy to use and
understand while still providing good performance.
[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReporCard-URL]
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FUHN%2Fggql.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FUHN%2Fggql?ref=badge_shield)
## Features

 - Simple binding of GraphQL schema elements to golang types and functions.
 - Multiple resolver options including dynamic resolvers.
 - High performance
 - No external dependencies.

## News

- The first release is out! Benchmarks look good, documentation is
  complete, and examples available.

## Using

The examples provide the best explanation for how to use the package.

 - [Examples](examples/README.md) of each resolver approach.
   - [Reflection Resolver Example](examples/reflection/README.md)
   - [Interface Resolver Example](examples/interface/README.md)
   - [Root Resolver Example](examples/root/README.md)

## Installation

```
go get github.com/uhn/ggql
```

To build and install the `ggqlgen` application:

```
go install ./...
```

## Releases

See [CHANGELOG.md](CHANGELOG.md)

## Benchmarks

Benchmarks are against
[graphql-go](https://github.com/graphql-go/graphql) which is currently
the most common golang GraphQL package. The ggql-i package is using
the interface resolvers (Resolver and ListResolver) while the ggql
framework is using reflection. For a more comprehensive comparison go
to
[graphql-benchmarks](https://github.com/the-benchmarker/graphql-benchmarks).

### Summary

| Rate                | Latency             | Verbosity       |
| ------------------- | ------------------- | --------------- |
| ggql-i (go)         | ggql-i (go)         | ggql (go)       |
| ggql (go)           | ggql (go)           | ggql-i (go)     |
| graphql-go (go)     | graphql-go (go)     | graphql-go (go) |

#### Parameters
- Last updated: 2019-11-18
- OS: Linux (version: 5.3.10-050310-generic, arch: x86_64)
- CPU Cores: 12
- Connections: 1000
- Duration: 20 seconds
- Units:
  - Rates are in requests per second.
  - Latency is in milliseconds.
  - Verbosity is the number of non-blank lines of code excluding comments.

### Rates
| Language  | Framework          |       Rate | Latency | Verbosity | Throughput |
|:----------|:-------------------|-----------:|--------:|----------:|-----------:|
| go (1.13) | ggql-i (0.9.9)     | **180487** |   0.067 |       253 |      26.10 |
| go (1.13) | ggql (0.9.9)       | **174742** |     067 |       176 |      25.25 |
| go (1.13) | graphql-go (0.7.8) |  **29614** |   0.086 |       378 |       4.28 |

### Latency
| Language  | Framework          |   Rate |   Latency | Verbosity | Average | 90th % | 99th % | Std Dev |
|:----------|:-------------------|-------:|----------:|----------:|--------:|-------:|-------:|--------:|
| go (1.13) | ggql-i (0.9.9)     | 180487 | **0.067** |       253 |   0.065 |  0.074 |  0.090 |    0.02 |
| go (1.13) | ggql (0.9.9)       | 174742 | **0.067** |       176 |   0.064 |  0.072 |  0.077 |    0.02 |
| go (1.13) | graphql-go (0.7.8) |  29614 | **0.086** |       378 |   0.085 |  0.091 |  0.099 |    0.02 |

### Verbosity
| Language  | Framework          |   Rate | Latency | Verbosity |
|:----------|:-------------------|-------:|--------:|----------:|
| go (1.13) | ggql (0.9.9)       | 174742 |   0.067 |   **176** |
| go (1.13) | ggql-i (0.9.9)     | 180487 |   0.067 |   **253** |
| go (1.13) | graphql-go (0.7.8) |  29614 |   0.086 |   **378** |

## More Information

 - [Overview](overview.md) of the package and how to use it.
 - [FAQ](faq.md)
 - [Go Docs](https://uhn.github.io/ggql)


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FUHN%2Fggql.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FUHN%2Fggql?ref=badge_large)