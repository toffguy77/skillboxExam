package common

import (
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"strings"
)

type ValidateData interface {
	models.SMSData | models.MMSData | models.EmailData
	HasCountry(map[string]models.Country) bool
}

type EmptyData interface {
	models.SMSData | models.MMSData | models.EmailData
	HasCountry(map[string]models.Country) bool
}

type ParseData interface {
	models.SMSData | models.MMSData | models.EmailData
	HasCountry(map[string]models.Country) bool
}

func IsTrustedProvider(provider string) bool {
	switch strings.ToLower(provider) {
	case
		"topolo", "rond", "kildy":
		return true
	}
	log.Printf("prepare provider is not trusted: %s\n", provider)
	return false
}

func Validate[T ValidateData](data []T, country map[string]models.Country) []T {
	for i := 0; i < len(data); i++ {
		if i == 0 && !data[i].HasCountry(country) {
			log.Printf("prepare data is not valid: %v\n", data[i])
			data = data[i+1:]
			i--
		}
		if !data[i].HasCountry(country) {
			log.Printf("prepare data is not valid: %v\n", data[i])
			data = append(data[:i-1], data[i+1:]...)
			i--
		}
	}
	return data
}

func IsEmptyData[T EmptyData](data T) bool {
	var copyData interface{}
	copyData = data
	if _, ok := copyData.(models.SMSData); ok {
		if copyData == (models.SMSData{}) {
			return true
		}
		if copyData == (models.MMSData{}) {
			return true
		}
		if copyData == (models.EmailData{}) {
			return true
		}
	}
	return false
}

func IsCorrectLine(line []string, l int) bool {
	if len(line) != l {
		return false
	}
	for _, val := range line {
		if val == "" {
			return false
		}
	}
	return true
}
