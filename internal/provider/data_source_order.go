package provider

import (
	"context"
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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"total": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modified": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhook": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lineitems": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pagecount": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"to": &schema.Schema{
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"address1": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"address2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"city": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"postcode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"formatted": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"organization": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"from": &schema.Schema{
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"address1": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"address2": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"city": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"postcode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"country": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"formatted": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"organization": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"simplex": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"color": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"pricing": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"account": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_reference": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"channel": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cancelled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cancellation_reason": &schema.Schema{
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
	// todo: handle the case where the order does not exist and we gracefully SetID("") i guess.
	if err != nil {
		return diag.FromErr(err)
	}

	// maybe we want to d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	// so we always run and refresh the sweet, sweet datas
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
