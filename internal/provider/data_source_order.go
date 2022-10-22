package provider

import (
	"context"

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
				Type:     schema.TypeInt,
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

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}
