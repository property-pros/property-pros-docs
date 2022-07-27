package notePurchaseAgreement

import (
	"fmt"

	documents "github.com/vireocloud/property-pros-docs/documents"
)

type NotePurchaseAgreementPage struct {
	*documents.HtmlTemplateBase
	index uint
}

func NewNotePurchaseAgreementPage(index uint, content string, model interface{}) (*NotePurchaseAgreementPage, error) {
	page := &NotePurchaseAgreementPage{
		index: index,
	}

	base, err := documents.NewHtmlTemplateBase(fmt.Sprintf("notePurchaseAgreementPage%v", index), content, model)

	if err != nil {
		return nil, err
	}

	page.HtmlTemplateBase = base

	return page, nil
}

func (h *NotePurchaseAgreementPage) ToString() string {
	return h.HtmlTemplateBase.ToString()
}
