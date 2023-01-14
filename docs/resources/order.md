---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mailform_order Resource - terraform-provider-mailform"
subcategory: ""
description: |-
  Mailform order
---

# mailform_order (Resource)

Mailform order



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `from_address_1` (String) The street number and name of the sender of this envelope or postcard.
- `from_city` (String) The address city of the sender of this envelope or postcard.
- `from_country` (String) The address country of the sender of this envelope or postcard. Example "US"
- `from_name` (String) The name of the sender of this envelope or postcard.
- `from_postcode` (String) The address postcode or zip code of the sender of this envelope or postcard. Example "00000"
- `from_state` (String) The address state of the sender of this envelope or postcard. Example "WA"
- `service` (String) What shipping service/speed to use.
- `to_address_1` (String) The street number and name of the recipient of this envelope or postcard.
- `to_city` (String) The address state of the recipient of this envelope or postcard.
- `to_country` (String) The address country of the recipient of this envelope or postcard. Example "US"
- `to_name` (String) The name of the recipient of this envelope or postcard.
- `to_postcode` (String) The address postcode or zip code of the recipient of this envelope or postcard. Example "00000"
- `to_state` (String) The address postcode or zip code of the recipient of this envelope or postcard. Example "WA"

### Optional

- `amount` (Number) The amount of the check associated with this order, in cents. Required if a check is to be included in this order.
- `bank_account` (String) The identifier of the bank account for the check associated with this order. Required if a check is to be included in this order.
- `check_memo` (String) The memo line for the check associated with this order.
- `check_name` (String) The name of the recipient of the check associated with this order. Required if a check is to be included in this order.
- `check_number` (Number) The number of the check associated with this order. Required if a check is to be included in this order.
- `color` (Boolean) True if the document should be printed in color, false if the document should be printed in black and white.
- `company` (String) The company that this order should be associated with.
- `customer_reference` (String) An optional customer reference to be attached to the order.
- `flat` (Boolean) True if the document MUST be mailed in a flat envelope, false if it is acceptable to mail the document folded.
- `from_address_2` (String) The suite or room number of the sender of this envelope or postcard.
- `from_organization` (String) The organization or company associated with this address.
- `message` (String) The message to be printed on the non-picture side of a postcard..
- `pdf_file` (String) File path of PDF to be printed and mailed by mailform. Orders cannot be updated/deleted.
- `pdf_url` (String) URL of PDF to be printed and mailed by mailform.
- `simplex` (Boolean) True if the document should be printed one page to a sheet, false if the document can be printed on both sides of a sheet.
- `stamp` (Boolean) True if the document MUST use a real postage stamp, false if it is acceptable to mail the document using metered postage or an imprint.
- `to_address_2` (String) The suite or room number of the recipient of this envelope or postcard.
- `to_organization` (String) The organization or company associated with the recipient of this envelope or postcard.
- `webhook` (String) The webhook that should receive notifications about order updates to this order.

### Read-Only

- `account` (String)
- `cancellation_reason` (String)
- `cancelled` (String)
- `channel` (String)
- `created` (String)
- `id` (String) The ID of this resource.
- `lineitems` (List of Object) (see [below for nested schema](#nestedatt--lineitems))
- `modified` (String)
- `object` (String)
- `state` (String)
- `test_mode` (Boolean)
- `total` (Number)

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

