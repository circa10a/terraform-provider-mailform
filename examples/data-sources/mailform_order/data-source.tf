terraform {
  required_providers {
    mailform = {
      source = "circa10a/mailform"
    }
  }
}

data "mailform_order" "sample_order" {
  id = "foo"
}

output "order_info" {
  value = data.mailform_order.sample_order
}
