package controllers

import (
	"testing"

	propertyProsApi "github.com/vireocloud/property-pros-sdk/api/note_purchase_agreement/v1"
)

func TestGetNotePurchaseAgreementDoc(t *testing.T) {
	controller := &NotePurchaseAgreementController{}

	result, err := controller.GetNotePurchaseAgreementDoc(nil, &propertyProsApi.GetNotePurchaseAgreementDocRequest{
		Payload: &propertyProsApi.NotePurchaseAgreementRecord{
			FirstName: "John",
			LastName:  "smith",
		},
	})

	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Errorf("Expected controller.GetNotePurchaseAgreementDoc to not return nil")
	}
}
