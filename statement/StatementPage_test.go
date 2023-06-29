package statement

import (
	"strings"
	"testing"
)

var testTemplate = "<div>{{.Test}}</div>"

type templateModel struct {
	Test string
}

var testTemplateModel = &templateModel{Test: "test"}

func TestNewstatementPage(t *testing.T) {
	page1, err := NewstatementPage(0, testTemplate, testTemplateModel)

	if err != nil {
		t.Error(err)
	}

	if page1 == nil {
		t.Errorf("unexpected nil returned from NewstatementPage(0)")
	}

	if page1.index != 0 {
		t.Errorf("expected statementPage.index to equal 0;  equals %v", page1.index)
	}

	if page1.Name() != "statementPage0" {
		t.Errorf("expected statementPage.templatePath to equal statementPage0;  equals %v", page1.Name())
	}

	if page1.HtmlTemplateBase == nil {
		t.Errorf("expected statementPage.template to not equal nil")
	}

	if page1.Name() != "statementPage0" {
		t.Errorf("expected statementPage.template.Name() to return statementPage0;  equals %v", page1.Name())
	}

	result := page1.ToString()

	expectedTestResult := strings.Replace(testTemplate, "{{.Test}}", testTemplateModel.Test, -1)

	if result != expectedTestResult {
		t.Errorf("expected statementPage.template.Execute() to return %v;  equals %v", expectedTestResult, result)
	}
}
