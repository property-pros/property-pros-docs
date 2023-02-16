package documents

import (
	"bytes"
	"encoding/json"
	"errors"
	"html"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	interfaces "github.com/vireocloud/property-pros-docs/interfaces"
	"go.etcd.io/etcd/pkg/fileutil"
)

type Pdf struct {
	*HtmlTemplateBase
	interfaces.IDocument
	pdfGenerator *wkhtml.PDFGenerator
	content      string
}

func (p *Pdf) AddPage(reader io.Reader) error {

	buffer, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	contentFile, _ := p.savePageToTempFile(string(buffer))
	return p.addPageToInternalPdfManager(contentFile)
}

func (p *Pdf) savePageToTempFile(content string) (*os.File, error) {

	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exPath := filepath.Dir(ex)

	file, err := os.CreateTemp(exPath, "notepurchaseagreement*.html")

	if err != nil {
		return file, err
	}

	pageContent := &strings.Builder{}
	p.HtmlTemplateBase.template.Execute(pageContent, &PdfPageTemplate{
		Content: template.HTML(content),
	})

	file.WriteString(pageContent.String())

	return file, nil
}

func (p *Pdf) addPageToInternalPdfManager(file *os.File) error {

	page := wkhtml.NewPage(file.Name())

	page.DisableSmartShrinking.Set(false)

	page.EnableLocalFileAccess.Set(true)
	page.LoadErrorHandling.Set("ignore")

	p.pdfGenerator.AddPage(page)

	return nil
}

func (p *Pdf) prepContent() error {
	return nil
}

func (p *Pdf) SetPages(readers []io.Reader) {

	pages := []wkhtml.PageProvider{}

	for _, reader := range readers {
		page := wkhtml.NewPageReader(reader)

		page.DisableSmartShrinking.Set(false)
		page.EnableLocalFileAccess.Set(true)

		pages = append(pages, page)
	}

	p.pdfGenerator.SetPages(pages)
}

func (p *Pdf) GetFileContent() (io.Reader, error) {

	p.prepContent()
	// log.Println("contents: ", p.content)
	err := p.pdfGenerator.Create()

	return bytes.NewReader(p.pdfGenerator.Bytes()), err
}

func (p *Pdf) Copy() interfaces.IDocument {
	document := *p

	return &document
}

func (p *Pdf) SaveDocumentToFile(filePath string) error {

	p.prepContent()

	err := p.pdfGenerator.Create()

	if err != nil {
		return err
	}

	return p.pdfGenerator.WriteFile(filePath)
}

func (p *Pdf) GetHtml() (string, error) {
	p.prepContent()

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

	inputFile := pdfMap.Pages[0].InputFile

	if fileutil.Exist(inputFile) {
		fileContent, err := os.ReadFile(inputFile)

		return html.UnescapeString(string(fileContent)), err
	}

	return "", errors.New("failed to read pdf content")
}

func NewPdf(pdfGenerator *wkhtml.PDFGenerator, template *HtmlTemplateBase) (*Pdf, error) {

	pdfGenerator.PageSize.Set("letter")

	return &Pdf{
		HtmlTemplateBase: template,
		pdfGenerator:     pdfGenerator,
	}, nil
}

type PdfPageTemplate struct {
	Content template.HTML
}

type PdfModel struct {
	Pages []PdfPageModel
}

type PdfPageModel struct {
	InputFile string
}
