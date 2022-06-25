package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	documents "github.com/vireocloud/property-pros-docs/Documents"
	notepurchaseagreement "github.com/vireocloud/property-pros-docs/notePurchaseAgreement"
	propertyProsApi "github.com/vireocloud/property-pros-docs/proto"
)

type NotePurchaseAgreementController struct {
	propertyProsApi.UnsafeNotePurchaseAgreementServiceServer
}

func (c *NotePurchaseAgreementController) GetNotePurchaseAgreementDoc(ctx context.Context, req *propertyProsApi.GetNotePurchaseAgreementDocRequest) (*propertyProsApi.GetNotePurchaseAgreementDocResponse, error) {

	pagesFile, err := os.Open("../../NotePurchaseAgreement/content/pages.json")

	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(pagesFile)

	pages := []string{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &pages)

	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return nil, err
	}

	pdf := documents.NewPdf(pdfGenerator)
	//TODO: handle errors
	notePurchaseAgreementModel, err := notepurchaseagreement.NewNotePurchaseAgreement(pages, pdf)

	doc := notePurchaseAgreementModel.ToDoc().GetFileContent()

	buffer := &bytes.Buffer{}
	buffer.ReadFrom(doc)

	return &propertyProsApi.GetNotePurchaseAgreementDocResponse{
		FileContent: buffer.Bytes(),
	}, nil
}

var _ propertyProsApi.NotePurchaseAgreementServiceServer = (*NotePurchaseAgreementController)(nil)
