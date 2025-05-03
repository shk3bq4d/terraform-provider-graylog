# This repo
This is a fork of a fork of fork, etc..

Kudos to the greater than me people that came before me. My contribution is close to zero but
I have been an user for a long time and need to still to move to Graylog6.

While I am leaving the rest of the README.md as-is, some/most of it may not make sense at all for you.

The latest to have released the provider and from which I forked is one-2-one, but I have used predecessors as well in the past

# Terraform Provider: Graylog
This is a Terraform provider for managing resources within [Graylog](https://docs.graylog.org/).

## Getting Started
As this provider is published to the public [Terraform Registry](https://registry.terraform.io/providers/shk3bq4d/graylog),
you can install it like so (for Terraform 0.14+):
```hcl
provider "graylog" {
  web_endpoint_uri = "http://example.com/api"
  api_version      = "v3"
}

terraform {
  required_providers {
    graylog = {
      source  = "shk3bq4d/graylog"
    }
  }
}
```

For more detailed instructions and documentation on the resources and data sources supported, please go to
[Terraform Registry](https://registry.terraform.io/providers/shk3bq4d/graylog/latest/docs).

## Maintenance
This provider is maintained during free time, so if you are interested in helping to develop this further, you
are more than welcome to submit a pull request or raise a ticket if you'd prefer.

## Development

### Requirements
If you do wish to help develop this, you will need the following installed:
* [Go](http://www.golang.org) (see `go.mod` file for the correct version to install)
* [Go Linter](https://formulae.brew.sh/formula/golangci-lint)
* [GOPATH](http://golang.org/doc/code.html#GOPATH) (is correctly setup)
* [Terraform](https://www.terraform.io/downloads.html) (0.14+)

### Building
Simply run `make build`, and it will compile and create a binary, as well as print-out instructions
on how to configure Terraform to use this locally built provider.
```shell
$ make build
```

### Testing

#### Unit Tests
```shell
$ make test
```

### Acceptance Tests
```shell
$ make testacc
```

### Documentation
Every data source or resource added must have an accompanying docs page (see `docs` directory for examples).

Docs are written using Markdown, and you can use [this page](https://registry.terraform.io/tools/doc-preview) to preview what your docs will look like when rendered.
