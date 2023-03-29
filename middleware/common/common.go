package common

import (
	"github.com/toffguy77/statusPage/internal/countries"
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/middleware/prepare"
)

func GetResultData() models.ResultSetT {
	var result models.ResultSetT

	countriesList := countries.GetCountries()
	result.SMS = prepare.GetSmsData(countriesList)
	result.MMS = prepare.GetMmsData(countriesList)
	result.VoiceCall = prepare.GetVoicecallData(countriesList)
	result.Email = prepare.GetEmailData(countriesList)
	result.Billing = prepare.GetBillingData()
	result.Support = prepare.GetSupportData()
	result.Incidents = prepare.GetIncidentsData()

	return result
}

func CheckResults(results models.ResultSetT) bool {
	if results.MMS == nil {
		return false
	}
	emptyBilling := models.BillingData{
		CreateCustomer: false,
		Purchase:       false,
		Payout:         false,
		Recurring:      false,
		FraudControl:   false,
		CheckoutPage:   false,
	}
	if results.Billing == emptyBilling {
		return false
	}
	if results.Email == nil {
		return false
	}
	if results.SMS == nil {
		return false
	}
	if results.Incidents == nil {
		return false
	}
	if results.Support == nil {
		return false
	}
	if results.VoiceCall == nil {
		return false
	}
	return true
}
