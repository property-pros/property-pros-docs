package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"io/fs"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	documents "github.com/vireocloud/property-pros-docs/documents"
	notepurchaseagreement "github.com/vireocloud/property-pros-docs/notepurchaseagreement"
	notePurchaseAgreementPagesContent "github.com/vireocloud/property-pros-docs/notepurchaseagreement/content"
	statements "github.com/vireocloud/property-pros-docs/statement"
	statementPagesContent "github.com/vireocloud/property-pros-docs/statement/content"
	propertyProsApi "github.com/vireocloud/property-pros-sdk/api/note_purchase_agreement/v1"
	"github.com/vireocloud/property-pros-sdk/api/statement/v1"
)

type StatementController struct {
	statement.UnimplementedStatementServiceServer
}

func (c *StatementController) GetStatementDoc(ctx context.Context, req *statement.GetStatementDocRequest) (response *statement.GetStatementDocResponse, err error) {
	log.Println("GetStatementDoc: ", req.String())

	fsys := fs.FS(statementPagesContent.Content)

	pages, err := ReadPagesDirectory(fsys)

	if err != nil {
		return
	}

	pdfGenerator, err := CreatePDFGenerator()

	if err != nil {
		return
	}

	pageContainerTemplate, err := fs.ReadFile(fsys, "contentTemplate.html")

	if err != nil {
		return
	}

	pageContainerTemplateManager, err := documents.NewHtmlTemplateBase("statementpagecontainer", string(pageContainerTemplate), nil)

	if err != nil {
		return
	}

	pdf, err := documents.NewPdf(pdfGenerator, pageContainerTemplateManager)

	if err != nil {
		return
	}

	statementModel, err := statements.NewStatement(pages, pdf)

	if err != nil {
		return
	}

	statementModel = FillStatementModel(statementModel, req)

	doc, err := statementModel.ToDoc().GetFileContent()

	if err != nil {
		log.Println("failed to get file content: ", err)
		return
	}

	buffer := &bytes.Buffer{}

	buffer.ReadFrom(doc)

	statementModel.ToDoc().SaveDocumentToFile("./statementtest.pdf")

	response = &statement.GetStatementDocResponse{
		Document: buffer.Bytes(),
	}

	return response, nil
}

func FillStatementModel(model *statements.Statement, data *statement.GetStatementDocRequest) *statements.Statement {
	payload := data.GetPayload()

	model.Balance = payload.GetBalance()
	model.EmailAddress = payload.GetEmailAddress()
	model.EndPeriodDate = payload.GetEndPeriodDate()
	model.Id = payload.GetId()
	model.Password = payload.GetPassword()
	model.Principle = payload.GetPrinciple()
	model.StartPeriodDate = payload.GetStartPeriodDate()
	model.UserId = payload.GetUserId()
	model.TotalIncome = payload.GetTotalIncome()

	return model
}

type NotePurchaseAgreementController struct {
	propertyProsApi.UnimplementedNotePurchaseAgreementServiceServer
}

func (c *NotePurchaseAgreementController) GetNotePurchaseAgreementDoc(req *propertyProsApi.GetNotePurchaseAgreementDocRequest, stream propertyProsApi.NotePurchaseAgreementService_GetNotePurchaseAgreementDocServer) error {
	response := &propertyProsApi.GetNotePurchaseAgreementDocResponse{}

	fsys := fs.FS(notePurchaseAgreementPagesContent.Content)

	pages, err := ReadPagesJson(fsys)

	if err != nil {
		return err
	}

	pdfGenerator, err := CreatePDFGenerator()

	if err != nil {
		return err
	}

	pageContainerTemplate, err := fs.ReadFile(fsys, "contentTemplate.html")

	if err != nil {
		return err
	}

	pageContainerTemplateManager, err := documents.NewHtmlTemplateBase("pagecontainer", string(pageContainerTemplate), nil)

	if err != nil {
		return err
	}

	pdf, err := documents.NewPdf(pdfGenerator, pageContainerTemplateManager)

	if err != nil {
		return err
	}

	notePurchaseAgreementModel, err := notepurchaseagreement.NewNotePurchaseAgreement(pages, pdf)

	if err != nil {
		return err
	}

	log.Println("GetNotePurchaseAgreementDoc: ", req.String())
	requestCopy := *req
	FillNotePurchaseAgreementModel(notePurchaseAgreementModel, &requestCopy)

	doc, err := notePurchaseAgreementModel.ToDoc().GetFileContent()

	if err != nil {
		log.Println("failed to get file content: ", err)
		return err
	}

	buffer := &bytes.Buffer{}
	buffer.ReadFrom(doc)

	response.FileContent = buffer.Bytes()

	notePurchaseAgreementModel.ToDoc().SaveDocumentToFile("./test.pdf")

	fileContent := response.FileContent

	fmt.Println("chunking file content")
	chunkSize := 1024 * 1024 // 1MB chunks
	for i := 0; i < len(fileContent); i += chunkSize {
		end := i + chunkSize
		if end > len(fileContent) {
			end = len(fileContent)
		}
		chunk := fileContent[i:end]
		fmt.Println("sending chunk")
		if err := stream.Send(&propertyProsApi.GetNotePurchaseAgreementDocResponse{
			FileContent: chunk,
		}); err != nil {
			return err
		}
	}
	
	fmt.Println("done chunking")
	return nil
}

// ReadPagesJson reads the pages file from the file system and returns a slice of strings
func ReadPagesJson(fsys fs.FS) ([]string, error) {
	pagesFileContent, err := fs.ReadFile(fsys, "pages.json")

	if err != nil {
		return nil, err
	}

	pages := []string{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'pages' which we defined above
	json.Unmarshal(pagesFileContent, &pages)

	return pages, nil
}

// ReadPagesDirectory reads the pages directory from the file system and creates an array of strings from the contents of each file in the pages directory
func ReadPagesDirectory(fsys fs.FS) ([]string, error) {
	pagesDir, err := fs.ReadDir(fsys, "pages")

	if err != nil {
		return nil, err
	}

	pages := []string{}

	for _, page := range pagesDir {

		pageFileContent, err := fs.ReadFile(fsys, "pages/"+page.Name())

		if err != nil {
			return nil, err
		}

		pages = append(pages, string(pageFileContent))
	}

	return pages, nil
}

// CreatePDFGenerator creates a new PDF generator with the given options
func CreatePDFGenerator() (*wkhtmltopdf.PDFGenerator, error) {
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}
	pdfGenerator.Cover.EnableLocalFileAccess.Set(true)
	pdfGenerator.TOC.EnableLocalFileAccess.Set(true)

	return pdfGenerator, nil
}

// FillNotePurchaseAgreementModel fills the note purchase agreement model with the given data
func FillNotePurchaseAgreementModel(model *notepurchaseagreement.NotePurchaseAgreement, data *propertyProsApi.GetNotePurchaseAgreementDocRequest) {

	payload := data.GetPayload()

	model.FirstName = payload.GetFirstName()
	model.LastName = payload.GetLastName()
	model.DateOfBirth = payload.GetDateOfBirth()
	model.EmailAddress = payload.GetUser().GetEmailAddress()
	model.FundsCommitted = payload.GetFundsCommitted()
	model.HomeAddress = payload.GetHomeAddress()
	model.PhoneNumber = payload.GetPhoneNumber()
	model.SocialSecurity = payload.GetSocialSecurity()
}

var _ propertyProsApi.NotePurchaseAgreementServiceServer = (*NotePurchaseAgreementController)(nil)

// package controllers

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"log"

// 	"io/fs"

// 	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
// 	documents "github.com/vireocloud/property-pros-docs/documents"
// 	notepurchaseagreement "github.com/vireocloud/property-pros-docs/notepurchaseagreement"
// 	notePurchaseAgreementPagesContent "github.com/vireocloud/property-pros-docs/notepurchaseagreement/content"
// 	propertyProsApi "github.com/vireocloud/property-pros-sdk/api/note_purchase_agreement/v1"
// )

// type NotePurchaseAgreementController struct {
// 	propertyProsApi.UnimplementedNotePurchaseAgreementServiceServer
// }

// func (c *NotePurchaseAgreementController) GetNotePurchaseAgreementDoc(ctx context.Context, req *propertyProsApi.GetNotePurchaseAgreementDocRequest) (*propertyProsApi.GetNotePurchaseAgreementDocResponse, error) {
// 	log.Println("GetNotePurchaseAgreementDoc");
// 	response := &propertyProsApi.GetNotePurchaseAgreementDocResponse{}

// 	fsys := fs.FS(notePurchaseAgreementPagesContent.Content)
// 	pagesFileContent, err := fs.ReadFile(fsys, "pages.json")

// 	if err != nil {
// 		return nil, err
// 	}

// 	pages := []string{}

// 	// we unmarshal our byteArray which contains our
// 	// jsonFile's content into 'pages' which we defined above
// 	json.Unmarshal(pagesFileContent, &pages)

// 	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

// 	if err != nil {
// 		return nil, err
// 	}
// 	pdfGenerator.Cover.EnableLocalFileAccess.Set(true)
// 	pdfGenerator.TOC.EnableLocalFileAccess.Set(true)

// 	pageContainerTemplate, err := fs.ReadFile(fsys, "contentTemplate.html")

// 	if err != nil {
// 		return nil, err
// 	}

// 	pageContainerTemplateManager, err := documents.NewHtmlTemplateBase("pagecontainer", string(pageContainerTemplate), nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	pdf, err := documents.NewPdf(pdfGenerator, pageContainerTemplateManager)

// 	if err != nil {
// 		return nil, err
// 	}

// 	notePurchaseAgreementModel, err := notepurchaseagreement.NewNotePurchaseAgreement(pages, pdf)

// 	if err != nil {
// 		return response, err
// 	}

// 	notePurchaseAgreementModel.FirstName = req.Payload.FirstName
// 	notePurchaseAgreementModel.LastName = req.Payload.LastName
// 	notePurchaseAgreementModel.DateOfBirth = req.Payload.DateOfBirth
// 	notePurchaseAgreementModel.EmailAddress = req.Payload.User.EmailAddress
// 	notePurchaseAgreementModel.FundsCommitted = req.Payload.FundsCommitted
// 	notePurchaseAgreementModel.HomeAddress = req.Payload.HomeAddress
// 	notePurchaseAgreementModel.PhoneNumber = req.Payload.PhoneNumber
// 	notePurchaseAgreementModel.SocialSecurity = req.Payload.SocialSecurity

// 	doc, err := notePurchaseAgreementModel.ToDoc().GetFileContent()

// 	if err != nil {
// 		log.Println("failed to get file content: ", err)
// 		return response, err
// 	}

// 	buffer := &bytes.Buffer{}
// 	buffer.ReadFrom(doc)

// 	response.FileContent = buffer.Bytes()

// 	notePurchaseAgreementModel.ToDoc().SaveDocumentToFile("./test.pdf")

// 	return response, nil
// }

// var _ propertyProsApi.NotePurchaseAgreementServiceServer = (*NotePurchaseAgreementController)(nil)
