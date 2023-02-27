terraform {
  required_providers {
    mailform = {
      source = "circa10a/mailform"
    }
  }
}

resource "mailform_pdf" "example" {
  header   = "My Resume"
  content  = "Some resume contents"
  filename = "./test.pdf"
}

resource "mailform_pdf" "converted_image" {
  image_filename = "./myimage.jpg"
  filename       = "./myimage.pdf"
}
