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
Change the OS_ARCH= variable in the makefile before installing. By default it's setup to a 64bit linux machine.
```sh
$ make install
```

## Using the provider

Please see the [examples](examples) directory for an example on how to use each resource and data source.

The provider itself requires your CloudCraft API key to be specified (and has two further optional arguments). This can be done either as an argument in your `provider` block, or through an environment variables.

### Provider Arguments

| Argument Name | Environment Variable     | Type              | Description                               | Required? | Default             |
|---------------|--------------------------|-------------------|-------------------------------------------|-----------|---------------------|
| `apikey`      | `CLOUDCRAFT_APIKEY`      | String, Sensitive | API Key for CloudCraft                    | Yes       |                     |
| `baseurl`     | `CLOUDCRAFT_BASEURL`     | String            | Host URL for cloudcraft.                  | No        | `api.cloudcraft.co` |
| `max_retries` | `CLOUDCRAFT_MAX_RETRIES` | Number            | Max retries for each CloudCraft API call. | No        | `1`                 |

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.15+ is _required_). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.


Enter the provider directory and build the provider

```sh
$ cd terraform-provider-cloudcraft
$ make init
$ make build
```

Install the provider to the local terraform directories
Change the OS_ARCH= variable in the makefile before installing. By default it's setup to a 64bit linux machine.
```sh
$ make install
```
