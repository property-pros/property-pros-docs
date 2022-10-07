package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"io/fs"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	documents "github.com/vireocloud/property-pros-docs/documents"
	propertyProsApi "github.com/vireocloud/property-pros-docs/generated/notepurchaseagreement"
	notepurchaseagreement "github.com/vireocloud/property-pros-docs/notepurchaseagreement"
	pagesContent "github.com/vireocloud/property-pros-docs/notepurchaseagreement/content"
)

type NotePurchaseAgreementController struct {
	propertyProsApi.UnsafeNotePurchaseAgreementServiceServer
}

func (c *NotePurchaseAgreementController) GetNotePurchaseAgreementDoc(ctx context.Context, req *propertyProsApi.GetNotePurchaseAgreementDocRequest) (*propertyProsApi.GetNotePurchaseAgreementDocResponse, error) {
	response := &propertyProsApi.GetNotePurchaseAgreementDocResponse{}

	fsys := fs.FS(pagesContent.Content)
	pagesFileContent, err := fs.ReadFile(fsys, "pages.json")

	if err != nil {
		return nil, err
	}

	pages := []string{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'pages' which we defined above
	json.Unmarshal(pagesFileContent, &pages)

	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}
	pdfGenerator.Cover.EnableLocalFileAccess.Set(true)
	pdfGenerator.TOC.EnableLocalFileAccess.Set(true)

	pageContainerTemplate, err := fs.ReadFile(fsys, "contentTemplate.html")

	if err != nil {
		return nil, err
	}

	pageContainerTemplateManager, err := documents.NewHtmlTemplateBase("pagecontainer", string(pageContainerTemplate), nil)

	if err != nil {
		return nil, err
	}

	pdf, err := documents.NewPdf(pdfGenerator, pageContainerTemplateManager)

	if err != nil {
		return nil, err
	}

	notePurchaseAgreementModel, err := notepurchaseagreement.NewNotePurchaseAgreement(pages, pdf)

	if err != nil {
		return response, err
	}

	notePurchaseAgreementModel.FirstName = req.Payload.FirstName
	notePurchaseAgreementModel.LastName = req.Payload.LastName
	notePurchaseAgreementModel.DateOfBirth = req.Payload.DateOfBirth
	notePurchaseAgreementModel.EmailAddress = req.Payload.EmailAddress
	notePurchaseAgreementModel.FundsCommitted = req.Payload.FundsCommitted
	notePurchaseAgreementModel.HomeAddress = req.Payload.HomeAddress
	notePurchaseAgreementModel.PhoneNumber = req.Payload.PhoneNumber
	notePurchaseAgreementModel.SocialSecurity = req.Payload.SocialSecurity

	doc, err := notePurchaseAgreementModel.ToDoc().GetFileContent()

	if err != nil {
		log.Println("failed to get file content: ", err)
		return response, err
	}

	buffer := &bytes.Buffer{}
	buffer.ReadFrom(doc)

	response.FileContent = buffer.Bytes()

	notePurchaseAgreementModel.ToDoc().SaveDocumentToFile("./test.pdf")

	return response, nil
}

var _ propertyProsApi.NotePurchaseAgreementServiceServer = (*NotePurchaseAgreementController)(nil)
