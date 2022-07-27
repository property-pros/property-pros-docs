package notePurchaseAgreement

import (
	"log"
	"strings"

	interfaces "github.com/vireocloud/property-pros-docs/interfaces"
)

type NotePurchaseAgreement struct {
	pages          []*NotePurchaseAgreementPage
	document       interfaces.IDocument
	FirstName      string
	LastName       string
	DateOfBirth    string
	HomeAddress    string
	EmailAddress   string
	PhoneNumber    string
	SocialSecurity string
	FundsCommitted uint64
}

func (this *NotePurchaseAgreement) ToDoc() interfaces.IDocument {

	document := this.document.Copy()

	for _, page := range this.pages {
		log.Println("page")
		document.AddPage(strings.NewReader(page.ToString()))
	}

	return document
}

func NewNotePurchaseAgreement(pages []string, document interfaces.IDocument) (*NotePurchaseAgreement, error) {
	doc := &NotePurchaseAgreement{}

	for i, pageContent := range pages {
		page, err := NewNotePurchaseAgreementPage(uint(i), pageContent, doc)

		if err != nil {
			return nil, err
		}

		doc.pages = append(doc.pages, page)
	}

	doc.document = document

	return doc, nil

}
