package notepurchaseagreement

import "testing"

func TestNewNotePurchaseAgreement(t *testing.T) {
	testNotePurchaseAgreement, err := NewNotePurchaseAgreement([]string{"<div>{{.FirstName}}</div>"})

	testNotePurchaseAgreement.FirstName = "test"

	if err != nil {
		t.Error(err)
	}

	pageLength := len(testNotePurchaseAgreement.pages)

	if pageLength != 1 {
		t.Errorf("expected length of NotePurchaseAgreementModel.pages to equal 1; equals %v", pageLength)
	}

	testPage := testNotePurchaseAgreement.pages[0]
	if testPage == nil {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0] to not equal nil")
	}

	if testPage.index != 0 {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0].index to equal 0; equals %v", testPage.index)
	}

	if testPage.template.Name() != "notePurchaseAgreementPage0" {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0].template.Name() to equal notePurchaseAgreementPage0; equals %v", testPage.template.Name())
	}

	expectedTemplateResult := "<div>test</div>"
	actualTemplateResult := testPage.ToString()

	if actualTemplateResult != expectedTemplateResult {
		t.Errorf("expected NotePurchaseAgreementModel.pages[0].ToString() to equal %v; equals %v", expectedTemplateResult, actualTemplateResult)
	}
}
