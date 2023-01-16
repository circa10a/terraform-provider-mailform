terraform {
  required_providers {
    mailform = {
      source = "github.com/circa10a/terraform-provider-mailform"
    }
  }
}

data "mailform_order" "sample_order" {
  id = "foo"
}

output "order_info" {
  value = data.mailform_order.sample_order
}
