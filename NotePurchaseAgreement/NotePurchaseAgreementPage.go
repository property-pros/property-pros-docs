package notepurchaseagreement

import (
	"fmt"
)

type NotePurchaseAgreementPage struct {
	*HtmlTemplateBase
	index uint
}

func NewNotePurchaseAgreementPage(index uint, content string, model interface{}) (*NotePurchaseAgreementPage, error) {
	page := &NotePurchaseAgreementPage{
		index: index,
	}

	base, err := NewHtmlTemplateBase(fmt.Sprintf("notePurchaseAgreementPage%v", index), content, model)

	if err != nil {
		return nil, err
	}

	page.HtmlTemplateBase = base

	return page, nil
}

func (h *NotePurchaseAgreementPage) ToString() string {
	return h.HtmlTemplateBase.ToString()
}
