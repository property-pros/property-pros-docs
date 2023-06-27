package controllers

import (
	"context"
	"testing"

	auth "github.com/vireocloud/property-pros-sdk/api/auth/v1"
	propertyProsApi "github.com/vireocloud/property-pros-sdk/api/note_purchase_agreement/v1"
)

func TestGetNotePurchaseAgreementDoc(t *testing.T) {
	controller := &NotePurchaseAgreementController{}

	result, err := controller.GetNotePurchaseAgreementDoc(context.TODO(), &propertyProsApi.GetNotePurchaseAgreementDocRequest{
		Payload: &propertyProsApi.NotePurchaseAgreementRecord{
			FirstName: "John",
			LastName:  "smith",
			User: &auth.User{
				EmailAddress: "test@yahoo.com",
			},
		}, 
	})

	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Errorf("Expected controller.GetNotePurchaseAgreementDoc to not return nil")
	}
}
