package common

import (
	"github.com/toffguy77/statusPage/internal/models"
	"testing"
)

func TestCheckResults(t *testing.T) {
	type args struct {
		results models.ResultSetT
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support: []int{3, 135},
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: true,
		},
		{
			name: "empty sms data",
			args: args{models.ResultSetT{
				SMS: nil,
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support: []int{3, 135},
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: false,
		},
		{
			name: "empty mms data",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: nil,
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support: []int{3, 135},
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: false,
		},
		{
			name: "empty voicecall data",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: nil,
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support: []int{3, 135},
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: false,
		},
		{
			name: "empty email data",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: nil,
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support: []int{3, 135},
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: false,
		},
		{
			name: "empty billing data",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					CreateCustomer: false,
					Purchase:       false,
					Payout:         false,
					Recurring:      false,
					FraudControl:   false,
					CheckoutPage:   false,
				},
				Support: []int{3, 135},
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: false,
		},
		{
			name: "empty support data",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support: nil,
				Incidents: []models.IncidentData{{
					Topic:  "Checkout page is down",
					Status: "active",
				}},
			}},
			want: false,
		},
		{
			name: "empty incident data",
			args: args{models.ResultSetT{
				SMS: [][]models.SMSData{[]models.SMSData{{
					Country:      "Австрия",
					Bandwidth:    "1",
					ResponseTime: "499",
					Provider:     "Kildy"},
				}},
				MMS: [][]models.MMSData{[]models.MMSData{{
					Country:      "Австрия",
					Provider:     "Kildy",
					Bandwidth:    "85",
					ResponseTime: "911"},
				}},
				VoiceCall: []models.VoiceCallData{{
					Country:             "RU",
					Bandwidth:           "44",
					ResponseTime:        "669",
					Provider:            "TransparentCalls",
					ConnectionStability: 0.71,
					TTFB:                904,
					VoicePurity:         26,
					MedianOfCallsTime:   39},
				},
				Email: map[string][][]models.EmailData{
					"AT": [][]models.EmailData{[]models.EmailData{{
						Country:      "AT",
						Provider:     "GMX",
						DeliveryTime: 15,
					},
						{
							Country:      "AT",
							Provider:     "Hotmail",
							DeliveryTime: 21,
						}}},
				},
				Billing: models.BillingData{
					Purchase:       true,
					CreateCustomer: true,
					Payout:         false,
					Recurring:      false,
					FraudControl:   true,
					CheckoutPage:   true,
				},
				Support:   []int{3, 135},
				Incidents: nil,
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckResults(tt.args.results); got != tt.want {
				t.Errorf("CheckResults() = %v, want %v", got, tt.want)
			}
		})
	}
}
