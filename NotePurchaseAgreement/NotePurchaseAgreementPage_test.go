package notepurchaseagreement

import (
	"strings"
	"testing"
)

type templateModel struct {
	Test string
}

func TestNewNotePurchaseAgreementPage(t *testing.T) {
	page1, err := NewNotePurchaseAgreementPage(0, testTemplate, testTemplateModel)

	if err != nil {
		t.Error(err)
	}

	if page1 == nil {
		t.Errorf("unexpected nil returned from NewNotePurchaseAgreementPage(0)")
	}

	if page1.index != 0 {
		t.Errorf("expected NotePurchaseAgreementPage.index to equal 0;  equals %v", page1.index)
	}

	if page1.templatePath != "notePurchaseAgreementPage0" {
		t.Errorf("expected NotePurchaseAgreementPage.templatePath to equal notePurchaseAgreementPage0;  equals %v", page1.templatePath)
	}

	if page1.template == nil {
		t.Errorf("expected NotePurchaseAgreementPage.template to not equal nil")
	}

	if page1.template.Name() != "notePurchaseAgreementPage0" {
		t.Errorf("expected NotePurchaseAgreementPage.template.Name() to return notePurchaseAgreementPage0;  equals %v", page1.template.Name())
	}

	stringBuilder := &strings.Builder{}

	page1.template.Execute(stringBuilder, testTemplateModel)

	expectedTestResult := strings.Replace(testTemplate, "{{.Test}}", testTemplateModel.Test, -1)

	if stringBuilder.String() != expectedTestResult {
		t.Errorf("expected NotePurchaseAgreementPage.template.Execute() to return %v;  equals %v", expectedTestResult, stringBuilder.String())
	}
}
