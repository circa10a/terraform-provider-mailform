package provider

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jung-kurt/gofpdf"
)

func resourcePDF() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Render a PDF and write to a local file.",

		CreateContext: resourcePDFCreate,
		ReadContext:   resourcePDFRead,
		DeleteContext: resourcePDFDelete,

		Schema: map[string]*schema.Schema{
			"filename": {
				Description: "The path to the PDF file that will be created",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"header": {
				Description: "Header/title of PDF",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				ConflictsWith: []string{
					"image_filename",
				},
			},
			"content": {
				Description: "Content of PDF",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				ConflictsWith: []string{
					"image_filename",
				},
			},
			"image_filename": {
				Description: "The image file to be converted to a PDF. Typically used for postcards",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				ConflictsWith: []string{
					"header",
					"content",
				},
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					buf := make([]byte, 512)

					imageFilename := val.(string)
					file, err := os.Open(imageFilename)
					if err != nil {
						errs = append(errs, err)
						return warns, errs
					}

					defer file.Close()

					_, err = file.Read(buf)
					if err != nil {
						errs = append(errs, err)
						return warns, errs
					}

					contentType := http.DetectContentType(buf)

					if contentType != "image/png" && contentType != "image/jpeg" {
						errs = append(errs, errors.New("image file is not a valid image"))
						return warns, errs
					}

					return warns, errs
				},
			},
		},
	}
}

func resourcePDFCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var diags diag.Diagnostics

	// Used for generating pdfs and also converting pdf's to images
	filename := d.Get("filename").(string)

	// If an image, convert to pdf
	if imageFilename, ok := d.GetOk("image_filename"); ok {
		err := convertImage(imageFilename.(string), filename)
		if err != nil {
			defer resourcePDFDelete(ctx, d, filename)
			return diag.FromErr(err)
		}
	} else {
		// Generate content if not image
		header := d.Get("header").(string)
		content := d.Get("content").(string)

		err := renderPDF(header, content, filename)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	outputContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return diag.FromErr(err)
	}

	checksum := sha1.Sum([]byte(outputContent))
	d.SetId(hex.EncodeToString(checksum[:]))

	tflog.Trace(ctx, "created a pdf resource")

	return diags
}

func resourcePDFRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// If the output file doesn't exist, mark the resource for creation.
	outputPath := d.Get("filename").(string)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		d.SetId("")
		return nil
	}

	// Verify that the content of the destination file matches the content we
	// expect. Otherwise, the file might have been modified externally, and we
	// must reconcile.
	outputContent, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return diag.FromErr(err)
	}

	outputChecksum := sha1.Sum(outputContent)
	if hex.EncodeToString(outputChecksum[:]) != d.Id() {
		d.SetId("")
		return nil
	}

	return nil
}

func resourcePDFDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	os.Remove(d.Get("filename").(string))
	return nil
}

// renderPDF converts header + content to a pdf and writes to an output file
func renderPDF(header string, content string, outputFilePath string) error {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", gofpdf.PageSizeLetter, "")
	pdf.AddPage()
	pdf.SetTitle(header, false)
	pdf.SetFont("Arial", "B", 16)
	// Calculate width of title and position
	wd := pdf.GetStringWidth(header) + 6
	pdf.SetX((210 - wd) / 2)
	// Title
	pdf.CellFormat(wd, 9, header, "", 1, "C", false, 0, "")
	// Line break
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 11)
	pdf.SetAutoPageBreak(true, 2.00)
	// Write ze content
	pdf.Write(8, content)

	return pdf.OutputFileAndClose(outputFilePath)
}

// convertImage converts input image path to pdf file
func convertImage(inputFilePath, outputFilePath string) error {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", gofpdf.PageSizeLetter, "")
	pdf.AddPage()
	pdf.Image(inputFilePath, 0, 0, 240, 480, false, "", 0, "")

	err := pdf.OutputFileAndClose(outputFilePath)
	if err != nil {
		return err
	}

	return nil
}
