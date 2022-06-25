package documents

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	interfaces "github.com/vireocloud/property-pros-docs/Interfaces"
)

type Pdf struct {
	interfaces.IDocument
	pdfGenerator *wkhtml.PDFGenerator
}

func (p *Pdf) AddPage(reader io.Reader) {

	page := wkhtml.NewPageReader(reader)

	p.pdfGenerator.AddPage(page)
}

func (p *Pdf) GetFileContent() io.Reader {
	return bytes.NewReader(p.pdfGenerator.Bytes())
}

func (p *Pdf) SaveDocumentToFile(filePath string) error {

	err := p.pdfGenerator.Create()

	if err != nil {
		return err
	}

	return p.pdfGenerator.WriteFile(filePath)
}

func (p *Pdf) GetHtml() (string, error) {

	jsonData, err := p.pdfGenerator.ToJSON()

	if err != nil {
		return "", err
	}

	pdfMap := &PdfModel{}

	err = json.Unmarshal(jsonData, &pdfMap)

	if err != nil {
		return "", err
	}

	if pdfMap.Pages == nil {
		return "", nil
	}

	if len(pdfMap.Pages) == 0 {
		return "", nil
	}

	encodedHtml, err := base64.StdEncoding.DecodeString(pdfMap.Pages[0].Base64PageData)

	return string(encodedHtml), err
}

type PdfModel struct {
	Pages []PdfPageModel
}

type PdfPageModel struct {
	Base64PageData string
}

func NewPdf(pdfGenerator *wkhtmltopdf.PDFGenerator) *Pdf {
	return &Pdf{
		pdfGenerator: pdfGenerator,
	}
}
