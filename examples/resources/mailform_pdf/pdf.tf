terraform {
  required_providers {
    mailform = {
      source = "github.com/nwheeler-splunk/mailform"
    }
  }
}

resource "mailform_pdf" "example" {
  header   = "My Resume"
  content  = "Some resume contents"
  filename = "./test.pdf"
}
