package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	SourceData  sourceDataType  `json:"source_data"`
	SupportVars supportVarsType `json:"support_vars"`
	ServerURL   string          `json:"server_url"`
}

type sourceDataType struct {
	CountriesDataFile string `json:"countries_data_file"`
	BillingDataFile   string `json:"billing_data_file"`
	EmailDataFile     string `json:"email_data_file"`
	IncidentURL       string `json:"incident_url"`
	MMSURL            string `json:"mmsurl"`
	SMSDataFile       string `json:"sms_data_file"`
	SupportURL        string `json:"support_url"`
	VoiceCallDataFile string `json:"voice_call_data_file"`
}

type supportVarsType struct {
	THRESHOLD_LOW    int `json:"threshold___low"`
	THRESHOLD_MEDIUM int `json:"threshold___medium"`
	CAPACITY         int `json:"capacity"`
	SPECIALISTS      int `json:"specialists"`
}

var Conf Config

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot load config file: %v\n", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Conf)
	if err != nil {
		log.Fatalf("cannot decode config file: %v\n", err)
	}
}
