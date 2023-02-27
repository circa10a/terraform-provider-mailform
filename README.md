# terraform-provider-mailform <img src="https://i.imgur.com/fAS7XqO.png" height="5%" width="5%" align="left"/>

This provider enables you to send physical mail, driven by terraform, via https://mailform.io

Standing on the shoulders of giants enables us to yeet mail further.

![Build Status](https://github.com/circa10a/terraform-provider-mailform/workflows/release/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/terraform-provider-mailform)](https://goreportcard.com/report/github.com/circa10a/terraform-provider-mailform)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/circa10a/terraform-provider-mailform?style=plastic)
[![Buy Me A Coffee](https://img.shields.io/badge/BuyMeACoffee-Donate-ff813f.svg?logo=CoffeeScript&style=plastic)](https://www.buymeacoffee.com/caleblemoine)

> :warning: Orders cannot be updated/deleted (cancelled). Once created,no more modifications can be made due to API limitations. Deleted resources are simply removed from state.

## Usage

The provider with use the environment variable `MAILFORM_API_TOKEN` by default unless specified in the provider configuration.

```hcl
terraform {
  required_providers {
    mailform = {
      source = "circa10a/mailform"
    }
  }
}

provider "mailform" {
  api_token = "XXX" // If not specified, will read MAILFORM_API_TOKEN environment variable
}

// Create PDF
resource "mailform_pdf" "example" {
  header   = "My Resumes"
  content  = "Some resume contents"
  filename = "./test.pdf"
}

// Create mail order
resource "mailform_order" "example" {
  pdf_file       = mailform_pdf.example.filename
  service        = "USPS_PRIORITY"
  to_name        = "A name"
  to_address_1   = "Address 1"
  to_city        = "Seattle"
  to_state       = "WA"
  to_postcode    = "00000"
  to_country     = "US"
  from_name      = "My name"
  from_address_1 = "My Address 1"
  from_city      = "Dallas"
  from_state     = "TX"
  from_postcode  = "00000"
  from_country   = "US"
}

// Fetch order data
data "mailform_order" "example" {
  id = mailform_order.example.id
}

output "order_info" {
  value = data.mailform_order.example
}
```

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
- [Go](https://golang.org/doc/install) >= 1.18

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```sh
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```sh
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.
