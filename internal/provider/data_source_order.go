package provider

import (
	"context"
	"strings"
	"time"

	"github.com/circa10a/go-mailform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Order data source in the Terraform provider mailform.",

		ReadContext: dataSourceOrderRead,

		Schema: map[string]*schema.Schema{
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
						"to": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"address1": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"address2": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"city": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"postcode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"formatted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"organization": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"from": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"address1": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"address2": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"city": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"postcode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"formatted": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"organization": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
							Type:     schema.TypeBool,
							Computed: true,
						},
						"pricing": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
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
		},
	}
}

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	if err := d.Set("lineitems", order.Data.Lineitems); err != nil {
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
