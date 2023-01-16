terraform {
  required_providers {
    mailform = {
      source = "github.com/circa10a/terraform-provider-mailform"
    }
  }
}

resource "mailform_pdf" "example" {
  header   = "My Resume"
  content  = "Some resume contents"
  filename = "./test.pdf"
}
