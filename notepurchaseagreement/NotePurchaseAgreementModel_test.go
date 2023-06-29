package notepurchaseagreement

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	docs "github.com/vireocloud/property-pros-docs/documents"
	documents "github.com/vireocloud/property-pros-docs/documents"
	interfaces "github.com/vireocloud/property-pros-docs/interfaces"
)

type testDoc struct {
}

func (d *testDoc) AddPage(io.Reader) {}

func (d *testDoc) SaveDocumentToFile(string) error {
	return nil
}

func (d *testDoc) GetFileContent() (io.Reader, error) {
	return nil, nil
}

func TestNewNotePurchaseAgreement(t *testing.T) {

	Setup()
	testNotePurchaseAgreement, err := NewNotePurchaseAgreement([]string{"<div>{{.FirstName}}</div>"}, testPdf)

	testNotePurchaseAgreement.FirstName = "test"

	if err != nil {
		t.Error(err)
	}

	pageLength := len(testNotePurchaseAgreement.pages)

	if pageLength != 1 {
		t.Errorf("expected length of NotePurchaseAgreementModel.pages to equal 1; equals %v", pageLength)
	}

	testPage := testNotePurchaseAgreement.pages[0]
	if testPage == nil {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0] to not equal nil")
	}

	if testPage.index != 0 {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0].index to equal 0; equals %v", testPage.index)
	}

	if testPage.HtmlTemplateBase.Name() != "notePurchaseAgreementPage0" {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0].template.Name() to equal notePurchaseAgreementPage0; equals %v", testPage.HtmlTemplateBase.Name())
	}

	expectedTemplateResult := "<div>test</div>"
	actualTemplateResult := testPage.ToString()

	if actualTemplateResult != expectedTemplateResult {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0].ToString() to equal %v; equals %v", expectedTemplateResult, actualTemplateResult)
	}

	doc := testNotePurchaseAgreement.ToDoc().(*docs.Pdf)

	if doc == nil {
		t.Errorf("Expected testNotePurchaseAgreement.ToDoc() to not return nil")
	}

	docReader, err := doc.GetFileContent()

	if err != nil {
		t.Error(err)
	}

	if docReader == nil {
		t.Errorf("Expected IDocument.GetFileContent() to not return nil")
	}

	buffer := &bytes.Buffer{}
	buffer.ReadFrom(docReader)
	log.Printf("doc: %v '%v'", buffer.Len(), buffer.String())

	if buffer.Len() == 0 {
		t.Errorf("Expected IDocument.GetFileContent() to not return empty")
	}

	doc.SaveDocumentToFile("./test.pdf")

	Teardown()
}

func Setup() {

	var err error

	testPdfGenerator, err = wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		fmt.Printf("setup error: %v", err)
		panic(err)
	}

	if testPdfGenerator == nil {
		fmt.Printf("failed to create pdf generator; NewPDFGenerator returned nil")

	}

	testTemplateBase, err := documents.NewHtmlTemplateBase("test", "<!doctype html><html><body>{{.Content}}</body></html><P style='page-break-before: always'>", nil)

	if err != nil {
		fmt.Printf("setup error: %v", err)
		panic(err)
	}

	if testPdfGenerator == nil {
		fmt.Printf("failed to create pdf generator; NewPDFGenerator returned nil")
	}

	testPdf, err = docs.NewPdf(testPdfGenerator, testTemplateBase)

	if err != nil {
		fmt.Printf("setup error: %v", err)
		panic(err)
	}
}

func Teardown() {
	testPdfGenerator = nil
	testPdf = nil
}

var testPdf interfaces.IDocument
var testPdfGenerator *wkhtmltopdf.PDFGenerator

func setup() {}

func teardown() {}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
