package provider

import (
	"context"

	"github.com/circa10a/go-mailform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"golang.org/x/exp/maps"
)

var orderInputSchema = map[string]*schema.Schema{
	"pdf_file": {
		Description:   "File path of PDF to be printed and mailed by mailform. Orders cannot be updated/deleted.",
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"pdf_url"},
		ForceNew:      true,
	},
	"pdf_url": {
		Description:   "URL of PDF to be printed and mailed by mailform.",
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"pdf_file"},
		ForceNew:      true,
	},
	"customer_reference": {
		Description: "An optional customer reference to be attached to the order.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"service": {
		Description:  "What shipping service/speed to use.",
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validation.StringInSlice(mailform.ServiceCodes, false),
		ForceNew:     true,
	},
	"webhook": {
		Description:  "The webhook that should receive notifications about order updates to this order.",
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validation.IsURLWithHTTPorHTTPS,
		ForceNew:     true,
	},
	"company": {
		Description: "The company that this order should be associated with.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"simplex": {
		Description: "True if the document should be printed one page to a sheet, false if the document can be printed on both sides of a sheet.",
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
	},
	"color": {
		Description: "True if the document should be printed in color, false if the document should be printed in black and white.",
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
	},
	"flat": {
		Description: "True if the document MUST be mailed in a flat envelope, false if it is acceptable to mail the document folded.",
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
	},
	"stamp": {
		Description: "True if the document MUST use a real postage stamp, false if it is acceptable to mail the document using metered postage or an imprint.",
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
	},
	"message": {
		Description: "The message to be printed on the non-picture side of a postcard..",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"to_name": {
		Description: "The name of the recipient of this envelope or postcard.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"to_organization": {
		Description: "The organization or company associated with the recipient of this envelope or postcard.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"to_address_1": {
		Description: "The street number and name of the recipient of this envelope or postcard.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"to_address_2": {
		Description: "The suite or room number of the recipient of this envelope or postcard.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"to_city": {
		Description: "The address state of the recipient of this envelope or postcard.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"to_state": {
		Description: "The address postcode or zip code of the recipient of this envelope or postcard. Example \"WA\"",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"to_postcode": {
		Description: "The address postcode or zip code of the recipient of this envelope or postcard. Example \"00000\"",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"to_country": {
		Description: "The address country of the recipient of this envelope or postcard. Example \"US\"",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"from_name": {
		Description: "The name of the sender of this envelope or postcard.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"from_organization": {
		Description: "The organization or company associated with this address.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"from_address_1": {
		Description: "The street number and name of the sender of this envelope or postcard.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"from_address_2": {
		Description: "The suite or room number of the sender of this envelope or postcard.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"from_city": {
		Description: "The address city of the sender of this envelope or postcard.",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"from_state": {
		Description: "The address state of the sender of this envelope or postcard. Example \"WA\"",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"from_postcode": {
		Description: "The address postcode or zip code of the sender of this envelope or postcard. Example \"00000\"",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"from_country": {
		Description: "The address country of the sender of this envelope or postcard. Example \"US\"",
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
	},
	"bank_account": {
		Description: "The identifier of the bank account for the check associated with this order. Required if a check is to be included in this order.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"amount": {
		Description: "The amount of the check associated with this order, in cents. Required if a check is to be included in this order.",
		Type:        schema.TypeInt,
		Optional:    true,
		ForceNew:    true,
	},
	"check_name": {
		Description: "The name of the recipient of the check associated with this order. Required if a check is to be included in this order.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	"check_number": {
		Description: "The number of the check associated with this order. Required if a check is to be included in this order.",
		Type:        schema.TypeInt,
		Optional:    true,
		ForceNew:    true,
	},
	"check_memo": {
		Description: "The memo line for the check associated with this order.",
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
	},
	// Computed
	"id": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

// getOrderCreateSchema merges the input fields for the mailform_order resource and computed fields for a mailform_order data source
func getOrderCreateSchema() map[string]*schema.Schema {
	merged := map[string]*schema.Schema{}
	maps.Copy(merged, orderSchema)
	// We copy input last so that webhook is set as input and not computed
	maps.Copy(merged, orderInputSchema)
	return merged
}

func resourceMailformOrder() *schema.Resource {
	return &schema.Resource{
		Description:   "Mailform order",
		CreateContext: resourceMailformOrderCreate,
		ReadContext:   orderRead,
		DeleteContext: resourceMailformOrderDelete,
		Schema:        getOrderCreateSchema(),
	}
}

func resourceMailformOrderCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	providerConfig := m.(map[string]interface{})
	client := providerConfig["client"].(*mailform.Client)
	order := mailform.OrderInput{
		FilePath:          d.Get("pdf_file").(string),
		URL:               d.Get("pdf_url").(string),
		CustomerReference: d.Get("customer_reference").(string),
		Service:           d.Get("service").(string),
		Webhook:           d.Get("webhook").(string),
		Company:           d.Get("company").(string),
		Simplex:           d.Get("simplex").(bool),
		Color:             d.Get("color").(bool),
		Flat:              d.Get("flat").(bool),
		Stamp:             d.Get("stamp").(bool),
		Message:           d.Get("message").(string),
		ToName:            d.Get("to_name").(string),
		ToOrganization:    d.Get("to_organization").(string),
		ToAddress1:        d.Get("to_address_1").(string),
		ToAddress2:        d.Get("to_address_2").(string),
		ToCity:            d.Get("to_city").(string),
		ToState:           d.Get("to_state").(string),
		ToPostcode:        d.Get("to_postcode").(string),
		ToCountry:         d.Get("to_country").(string),
		FromName:          d.Get("from_name").(string),
		FromOrganization:  d.Get("from_organization").(string),
		FromAddress1:      d.Get("from_address_1").(string),
		FromAddress2:      d.Get("from_address_2").(string),
		FromCity:          d.Get("from_city").(string),
		FromState:         d.Get("from_state").(string),
		FromPostcode:      d.Get("from_postcode").(string),
		FromCountry:       d.Get("from_country").(string),
		BankAccount:       d.Get("bank_account").(string),
		Amount:            d.Get("amount").(int),
		CheckName:         d.Get("check_name").(string),
		CheckNumber:       d.Get("check_number").(int),
		CheckMemo:         d.Get("check_memo").(string),
	}

	result, err := client.CreateOrder(order)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(result.Data.ID)

	// Set computed fields in state. Saves alot of copy paste by just running an extra GET after creating the order
	return orderRead(ctx, d, m)
}

func resourceMailformOrderDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// API doesn't support deleting orders, we simply just remove from state
	d.SetId("")
	return nil
}
