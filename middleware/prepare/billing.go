package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/billing"
	"log"
)

func GetBillingData(c chan models.BillingData) {
	defer close(c)
	billingProvider := billing.BillingProvider{
		Name: "Billing Status",
	}
	billingRes, err := billingProvider.GetStatus()
	if err != nil {
		log.Printf("can't get Billing status: %v", err)
		c <- models.BillingData{}
	}
	c <- billingRes
}
