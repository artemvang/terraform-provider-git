# Terraform Provider for Git

![](https://github.com/artemvang/terraform-provider-git/workflows/test/badge.svg)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
-	[Go](https://golang.org/doc/install) >= 1.14

## Usage

```hcl
provider "git" {
  private_key = file("key.pem")
}

data "git_repository" "example" {
  url = var.repo_url
}

```

## Contributing

To build the provider:

```sh
$ go build
```

To test the provider:

```sh
$ go test -v ./...
```

To run all acceptance tests:

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ TF_ACC=1 go test -v ./...
```

To run a subset of acceptance tests:

```sh
$ TF_ACC=1 go test -v ./... -run=TestAccDataSourceGitRepository
```
