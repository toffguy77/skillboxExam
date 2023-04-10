package config

type sourceDataType struct {
	CountriesDataFile string
	BillingDataFile   string
	EmailDataFile     string
	IncidentURL       string
	MMSURL            string
	SMSDataFile       string
	SupportURL        string
	VoiceCallDataFile string
}

var (
	SourceData = sourceDataType{
		CountriesDataFile: "internal/countries/countries.csv",
		BillingDataFile:   "sourcedata/billing.data",
		EmailDataFile:     "sourcedata/email.data",
		IncidentURL:       "http://127.0.0.1:8383/accendent",
		MMSURL:            "http://127.0.0.1:8383/mms",
		SMSDataFile:       "sourcedata/sms.data",
		SupportURL:        "http://127.0.0.1:8383/support",
		VoiceCallDataFile: "sourcedata/voice.data",
	}
)

type supportVarsType struct {
	THRESHOLD_LOW    int
	THRESHOLD_MEDIUM int
	CAPACITY         int
	SPECIALISTS      int
}

var (
	SupportVars = supportVarsType{
		THRESHOLD_LOW:    9,
		THRESHOLD_MEDIUM: 16,
		CAPACITY:         18,
		SPECIALISTS:      7,
	}
)

var ServerURL string = "127.0.0.1:8888"
