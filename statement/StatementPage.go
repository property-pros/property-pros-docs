package statement

import (
	"fmt"

	documents "github.com/vireocloud/property-pros-docs/documents"
)

type statementPage struct {
	*documents.HtmlTemplateBase
	index uint
}

func NewstatementPage(index uint, content string, model interface{}) (*statementPage, error) {
	page := &statementPage{
		index: index,
	}

	base, err := documents.NewHtmlTemplateBase(fmt.Sprintf("statementPage%v", index), content, model)

	if err != nil {
		return nil, err
	}

	page.HtmlTemplateBase = base

	return page, nil
}

func (h *statementPage) ToString() string {
	return h.HtmlTemplateBase.ToString()
}
