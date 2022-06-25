package documents

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	interfaces "github.com/vireocloud/property-pros-docs/Interfaces"
)

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

	testPdf = NewPdf(testPdfGenerator)
}

func Teardown() {
	testPdfGenerator = nil
	testPdf = nil
}

func TestGetFileContent(t *testing.T) {
	Setup()

	testReader := strings.NewReader("<div>test</div>\n")

	testPdf.AddPage(testReader)
	testPdfGenerator.Create()
	pdfContentReader := testPdf.GetFileContent().(*bytes.Reader)

	if pdfContentReader == nil {
		t.Errorf("expected file content reader to not be nil")
	}

	if pdfContentReader.Len() == 0 {
		t.Errorf("Expected pdfReader.Len() to be greater than 0")
	}

	if pdfContentReader.Size() == 0 {
		t.Errorf("Expected pdfReader.Size() to be greater than 0")
	}
	buf := &bytes.Buffer{}

	_, err := buf.ReadFrom(pdfContentReader)

	if err != nil {
		t.Error(err)
	}

	pdfContent := buf.Bytes()

	if pdfContent == nil {
		t.Errorf("expected file content to not be nil")
	}
	if len(pdfContent) == 0 {
		t.Errorf("expected file content to not be empty")
	}

	Teardown()
}

func TestAddPage(t *testing.T) {
	Setup()

	testReader := strings.NewReader("<div>test</div>")

	testPdf.AddPage(testReader)

	pdfHtml, err := testPdf.(*Pdf).GetHtml()

	if err != nil {
		t.Error(err)
	}

	if pdfHtml == "" {
		t.Errorf("testPdf.(*Pdf).GetHtml() to not return empty string")
	}

	if pdfHtml != "<div>test</div>" {
		t.Errorf("ExpectedpdfMap.Pages[0].Base64PageDat to decode to <div>test</div>;  decodes to %v", pdfHtml)
	}

	Teardown()
}

func TestSaveDocumentToFile(t *testing.T) {
	Setup()

	testReader := strings.NewReader("<div>test</div>")

	testPdf.AddPage(testReader)
	err := testPdf.SaveDocumentToFile("test.pdf")

	if err != nil {
		t.Errorf("Failed to save document to file;  error: %v", err)
	}

	file, err := os.Stat("test.pdf")

	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Failed to create pdf file; Error: %v", err)
	}

	if file.Size() == 0 {
		t.Errorf("Expected file to not be empty")
	}

	Teardown()
}

func TestNewPdf(t *testing.T) {
	Setup()

	if testPdf == nil {
		t.Errorf("failed to create instance of Pdf with NewPdf()")
	}

	Teardown()
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
