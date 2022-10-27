terraform {
  required_providers {
    mailform = {
      source = "github.com/nwheeler-splunk/mailform"
    }
  }
}

data "mailform_order" "sample_order" {
  id = "foo"
}

output "order_info" {
  value = data.mailform_order.sample_order
}
