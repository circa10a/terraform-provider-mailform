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

resource "mailform_order" "example" {
  pdf_file       = mailform_pdf.example.filename
  service        = "USPS_PRIORITY"
  to_name        = "A name"
  to_address_1   = "Address 1"
  to_city        = "Seattle"
  to_state       = "WA"
  to_postcode    = "00000"
  to_country     = "US"
  from_name      = "My name"
  from_address_1 = "My Address 1"
  from_city      = "Dallas"
  from_state     = "TX"
  from_postcode  = "00000"
  from_country   = "US"
}
