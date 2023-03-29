package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/billing"
	"log"
)

func GetBillingData() models.BillingData {
	billingProvider := billing.BillingProvider{
		Name: "Billing Status",
	}
	billingRes, err := billingProvider.GetStatus()
	if err != nil {
		log.Printf("can't get Billing status: %v", err)
		return models.BillingData{
			CreateCustomer: false,
			Purchase:       false,
			Payout:         false,
			Recurring:      false,
			FraudControl:   false,
			CheckoutPage:   false,
		}
	}
	return billingRes
}
