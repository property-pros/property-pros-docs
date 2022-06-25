package notepurchaseagreement

import (
	"strings"

	interfaces "github.com/vireocloud/property-pros-docs/Interfaces"
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

func (n *NotePurchaseAgreement) ToDoc() interfaces.IDocument {
	for _, page := range n.pages {
		n.document.AddPage(strings.NewReader(page.ToString()))
	}

	return n.document
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
