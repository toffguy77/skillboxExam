package models

import "reflect"

type Country struct {
	Name       string
	Alpha2     string
	Alpha3     string
	ISO_3166_1 string
	ISO_3166_2 string
}

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

type ResultT struct {
	Status bool       `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Data   ResultSetT `json:"data"`   // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки  (детали ниже)
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"prepare"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

func (m MMSData) HasCountry(country map[string]Country) bool {
	value := reflect.ValueOf(m)
	hasCountryFiled := value.FieldByName("Country")
	if !hasCountryFiled.IsValid() {
		return false
	}
	_, ok := country[m.Country]
	if ok {
		return true
	}
	return false
}

func (m SMSData) HasCountry(country map[string]Country) bool {
	value := reflect.ValueOf(m)
	hasCountryFiled := value.FieldByName("Country")
	if !hasCountryFiled.IsValid() {
		return false
	}
	_, ok := country[m.Country]
	if ok {
		return true
	}
	return false
}

func (m EmailData) HasCountry(country map[string]Country) bool {
	value := reflect.ValueOf(m)
	hasCountryFiled := value.FieldByName("Country")
	if !hasCountryFiled.IsValid() {
		return false
	}
	_, ok := country[m.Country]
	if ok {
		return true
	}
	return false
}

func (m VoiceCallData) HasCountry(country map[string]Country) bool {
	value := reflect.ValueOf(m)
	hasCountryFiled := value.FieldByName("Country")
	if !hasCountryFiled.IsValid() {
		return false
	}
	_, ok := country[m.Country]
	if ok {
		return true
	}
	return false
}
