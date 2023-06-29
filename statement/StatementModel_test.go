package statement

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
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

func TestNewstatement(t *testing.T) {

	Setup()
	pageContent := "<!doctype html><html><body><div>{{.FirstName}}</div></body></html><P style='page-break-before: always'>"
	teststatement, err := NewStatement([]string{pageContent}, testPdf)

	if err != nil {
		t.Error(err)
	}

	teststatement.FirstName = "test"

	pageLength := len(teststatement.pages)

	if pageLength != 1 {
		t.Errorf("expected length of statementModel.pages to equal 1; equals %v", pageLength)
	}

	testPage := teststatement.pages[0]
	if testPage == nil {
		t.Errorf("expected statementModel.pages[0] to not equal nil")
	}

	if testPage.index != 0 {
		t.Errorf("expected statementModel.pages[0].index to equal 0; equals %v", testPage.index)
	}

	if testPage.HtmlTemplateBase.Name() != "statementPage0" {
		t.Errorf("expected statementModel.pages[0].template.Name() to equal statementPage0; equals %v", testPage.HtmlTemplateBase.Name())
	}

	expectedTemplateResult := strings.Replace(pageContent, "{{.FirstName}}", teststatement.FirstName, -1)
	actualTemplateResult := testPage.ToString()

	if actualTemplateResult != expectedTemplateResult {
		t.Errorf("expected statementModel.pages[0].ToString() to equal %v; equals %v", expectedTemplateResult, actualTemplateResult)
	}

	doc := teststatement.ToDoc().(*docs.Pdf)

	if doc == nil {
		t.Errorf("Expected teststatement.ToDoc() to not return nil")
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
	// log.Printf("doc: %v '%v'", buffer.Len(), buffer.String())

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
	testPdfGenerator.Cover.EnableLocalFileAccess.Set(true)
	testPdfGenerator.TOC.EnableLocalFileAccess.Set(true)

	if testPdfGenerator == nil {
		fmt.Printf("failed to create pdf generator; NewPDFGenerator returned nil")

	}

	testTemplateBase, err := documents.NewHtmlTemplateBase("test", "<!doctype html><html><body>{{.Content}}</body></html><P style='page-break-before: always'>", nil)

	if err != nil {
		fmt.Printf("setup error: %v", err)
		panic(err)
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
