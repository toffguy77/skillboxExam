package common

import (
	"github.com/toffguy77/statusPage/internal/countries"
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/middleware/prepare"
)

func GetResultData() models.ResultSetT {
	var result models.ResultSetT

	countriesList := countries.GetCountries()
	chanSMS := make(chan [][]models.SMSData)
	chanMMS := make(chan [][]models.MMSData)
	chanVoiceCall := make(chan []models.VoiceCallData)
	chanEmail := make(chan map[string][][]models.EmailData)
	chanBilling := make(chan models.BillingData)
	chanSupport := make(chan []int)
	chanIncidents := make(chan []models.IncidentData)

	for {
		go prepare.GetSmsData(chanSMS, countriesList)
		go prepare.GetMmsData(chanMMS, countriesList)
		go prepare.GetVoiceCallData(chanVoiceCall, countriesList)
		go prepare.GetEmailData(chanEmail, countriesList)
		go prepare.GetBillingData(chanBilling)
		go prepare.GetSupportData(chanSupport)
		go prepare.GetIncidentsData(chanIncidents)
		break
	}

	result.SMS = <-chanSMS
	result.MMS = <-chanMMS
	result.VoiceCall = <-chanVoiceCall
	result.Email = <-chanEmail
	result.Billing = <-chanBilling
	result.Support = <-chanSupport
	result.Incidents = <-chanIncidents

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
