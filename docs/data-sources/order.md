---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mailform_order Data Source - terraform-provider-mailform"
subcategory: ""
description: |-
  Order data source in the Terraform provider mailform.
---

# mailform_order (Data Source)

Order data source in the Terraform provider mailform.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `account` (String)
- `cancellation_reason` (String)
- `cancelled` (String)
- `channel` (String)
- `created` (String)
- `customer_reference` (String)
- `id` (String) The ID of this resource.
- `lineitems` (List of Object) (see [below for nested schema](#nestedatt--lineitems))
- `modified` (String)
- `object` (String)
- `state` (String)
- `test_mode` (Boolean)
- `total` (Number)
- `webhook` (String)

<a id="nestedatt--lineitems"></a>
### Nested Schema for `lineitems`

Read-Only:

- `color` (Boolean)
- `from_address_1` (String)
- `from_address_2` (String)
- `from_city` (String)
- `from_country` (String)
- `from_formatted` (String)
- `from_name` (String)
- `from_organization` (String)
- `from_postcode` (String)
- `from_state` (String)
- `id` (String)
- `pagecount` (Number)
- `service` (String)
- `simplex` (Boolean)
- `to_address_1` (String)
- `to_address_2` (String)
- `to_city` (String)
- `to_country` (String)
- `to_formatted` (String)
- `to_name` (String)
- `to_organization` (String)
- `to_postcode` (String)
- `to_state` (String)


