# Terraform Provider for CloudCraft

## Create By
https://www.kotechnologies.co.uk

## CloudCraft
https://www.cloudcraft.co/

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x or above
- [Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

## Building The Provider

Clone repository to your host machine

```sh
$ git clone https://github.com/KOTechnologiesLtd/terraform-provider-cloudcraft
```

Enter the provider directory and build the provider

```sh
$ cd terraform-provider-cloudcraft
$ make init
$ make build
```

Install the provider to the local terraform directories

```sh
$ make install
```

## Using the provider

Please see the [examples](examples) directory for an example on how to use each.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.15+ is _required_). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.


Enter the provider directory and build the provider

```sh
$ cd terraform-provider-cloudcraft
$ make init
$ make build
```

Install the provider to the local terraform directories

```sh
$ make install
```
