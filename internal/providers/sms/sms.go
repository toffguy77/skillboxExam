package sms

import (
	"bufio"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"os"
	"strings"
)

const SMSDataFile = "sms.data"

//FIXME: fix SMS source

type SMSProvider struct {
	Name string
}

func (p SMSProvider) GetStatus(countries map[string]models.Country) ([]models.SMSData, error) {
	data, err := parseSmsData(SMSDataFile)
	if err != nil {
		log.Printf("can't parse prepare data: %v\n", err)
		return nil, err
	}
	result := validate(data, countries)
	return result, nil
}

func validate(data []models.SMSData, country map[string]models.Country) []models.SMSData {
	for iter, res := range data {
		_, ok := country[res.Country]
		if !ok {
			log.Printf("prepare data is not valid: %v\n", res)
			data[iter] = data[len(data)-1]
			data = data[:len(data)-1]
		}
	}
	return data
}

func parseSmsData(file string) ([]models.SMSData, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("error opening %s file: %v\n", file, err)
		return nil, err
	}
	defer f.Close()

	var data []models.SMSData
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dataLine := parseString(scanner.Text())
		if !isEmptySmsData(dataLine) {
			data = append(data, dataLine)
		}
	}
	return data, nil
}

func isEmptySmsData(data models.SMSData) bool {
	if data == (models.SMSData{}) {
		return true
	}
	return false
}

func parseString(line string) models.SMSData {
	parsedLine := strings.Split(line, ";")
	if !isCorrectLine(parsedLine, 4) {
		log.Printf("line is not valid: %s\n", line)
		return models.SMSData{}
	}
	if !isTrustedSmsProvider(parsedLine[3]) {
		log.Printf("untrusted provider, skip: %s\n", line)
		return models.SMSData{}
	}
	data := models.SMSData{
		Country:      parsedLine[0],
		Bandwidth:    parsedLine[1],
		ResponseTime: parsedLine[2],
		Provider:     parsedLine[3],
	}
	return data
}

func isCorrectLine(line []string, l int) bool {
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

func isTrustedSmsProvider(provider string) bool {
	switch strings.ToLower(provider) {
	case
		"topolo", "rond", "kildy":
		return true
	}
	log.Printf("prepare provider is not trusted: %s\n", provider)
	return false
}
