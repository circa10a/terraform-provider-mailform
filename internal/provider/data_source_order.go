package provider

import (
	"context"
	"strings"
	"time"

	"github.com/circa10a/go-mailform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var orderSchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"object": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"created": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"total": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"modified": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"webhook": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"lineitems": {
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"pagecount": {
					Type:     schema.TypeInt,
					Computed: true,
				},
				"to_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_address_1": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_address_2": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_city": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_postcode": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_country": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_formatted": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"to_organization": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_address_1": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_address_2": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_city": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_postcode": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_country": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_formatted": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"from_organization": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"simplex": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"color": {
					Type:     schema.TypeBool,
					Computed: true,
				},
				"service": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	},
	"account": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"customer_reference": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"channel": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"test_mode": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"state": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cancelled": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"cancellation_reason": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Order data source in the Terraform provider mailform.",
		ReadContext: orderRead,
		Schema:      orderSchema,
	}
}

func orderRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	providerConfig := m.(map[string]interface{})
	client := providerConfig["client"].(*mailform.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)
	order, err := client.GetOrder(id)
	if err != nil {
		// handle the case where the order does not exist and we gracefully SetID("") I guess.
		// this allows the user to make decisions in tf code instead of having that shit just bail out.
		if strings.Contains(err.Error(), "order_not_found") {
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}

	d.SetId(order.Data.ID)

	if err := d.Set("object", order.Data.Object); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("created", order.Data.Created.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("total", order.Data.Total); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("modified", order.Data.Modified.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("webhook", order.Data.Webhook); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("lineitems", flattenLineItems(order)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("account", order.Data.Account); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("customer_reference", order.Data.CustomerReference); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("channel", order.Data.Channel); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("test_mode", order.Data.TestMode); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", order.Data.State); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cancelled", order.Data.Cancelled.Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cancellation_reason", order.Data.CancellationReason); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func flattenLineItems(order *mailform.Order) []any {
	if order != nil {
		lineItems := order.Data.Lineitems
		ois := make([]interface{}, len(lineItems), len(lineItems))

		for i, orderItem := range lineItems {
			oi := make(map[string]interface{})

			oi["id"] = orderItem.ID
			oi["pagecount"] = orderItem.Pagecount
			oi["simplex"] = orderItem.Simplex
			oi["color"] = orderItem.Color
			oi["service"] = orderItem.Service
			oi["to_name"] = orderItem.To.Name
			oi["to_address_1"] = orderItem.To.Address1
			oi["to_address_2"] = orderItem.To.Address2
			oi["to_city"] = orderItem.To.City
			oi["to_postcode"] = orderItem.To.Postcode
			oi["to_country"] = orderItem.To.Country
			oi["to_formatted"] = orderItem.To.Formatted
			oi["to_organization"] = orderItem.To.Organization
			oi["from_name"] = orderItem.From.Name
			oi["from_address_1"] = orderItem.From.Address1
			oi["from_address_2"] = orderItem.From.Address2
			oi["from_state"] = orderItem.From.State
			oi["from_postcode"] = orderItem.From.Postcode
			oi["from_country"] = orderItem.From.Country
			oi["from_formatted"] = orderItem.From.Formatted
			oi["from_organization"] = orderItem.From.Organization

			ois[i] = oi
		}

		return ois
	}

	return make([]interface{}, 0)
}
