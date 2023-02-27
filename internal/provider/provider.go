package provider

import (
	"context"

	"github.com/circa10a/go-mailform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	mailformTokenAPIEnvVar = "MAILFORM_API_TOKEN"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_token": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(mailformTokenAPIEnvVar, nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"mailform_order": dataSourceOrder(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"mailform_order": resourceMailformOrder(),
				"mailform_pdf":   resourcePDF(),
			},
			ConfigureContextFunc: providerConfigure,
		}

		return p
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	api_token := d.Get("api_token").(string)
	client, err := mailform.New(&mailform.Config{
		Token: api_token,
	})
	if err != nil {
		return nil, diag.FromErr(err)
	}
	providerConfig := make(map[string]interface{})
	providerConfig["client"] = client
	return providerConfig, diags
}
