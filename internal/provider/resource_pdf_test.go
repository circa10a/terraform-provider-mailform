package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourcePDF(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePDF,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"mailform_pdf.example", "header", regexp.MustCompile("My Resume"),
					),

					resource.TestMatchResourceAttr(
						"mailform_pdf.example", "content", regexp.MustCompile("Some resume contents"),
					),
					resource.TestMatchResourceAttr(
						"mailform_pdf.example", "filename", regexp.MustCompile("./test.pdf"),
					),
				),
			},
		},
		CheckDestroy: checkFileDeleted("./test.pdf"),
	})
}

const testAccResourcePDF = `
resource "mailform_pdf" "example" {
	header   = "My Resume"
	content  = "Some resume contents"
	filename = "./test.pdf"
  }
`

func checkFileDeleted(shouldNotExistFile string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		if _, err := os.Stat(shouldNotExistFile); os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("file %s was not deleted", shouldNotExistFile)
	}
}
