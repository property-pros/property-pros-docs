package statement

import (
	"strings"

	interfaces "github.com/vireocloud/property-pros-docs/interfaces"
)

type Statement struct {
	pages       []*statementPage 
	document    interfaces.IDocument
	Id              string 
	UserId          string 
	EmailAddress    string 
	Password        string 
	StartPeriodDate string 
	EndPeriodDate   string 
	Balance         string
	TotalIncome     string 
	Principle       string 
}

func (s *Statement) ToDoc() interfaces.IDocument {

	document := s.document.Copy()

	for _, page := range s.pages {
		document.AddPage(strings.NewReader(page.ToString()))
	}

	return document
}

func NewStatement(pages []string, document interfaces.IDocument) (*Statement, error) {
	statement := &Statement{}

	for i, pageContent := range pages {
		page, err := NewstatementPage(uint(i), pageContent, statement)

		if err != nil {
			return nil, err
		}

		statement.pages = append(statement.pages, page)
	}

	statement.document = document

	return statement, nil

}
