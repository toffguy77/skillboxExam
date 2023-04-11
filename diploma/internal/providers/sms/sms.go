package sms

import (
	"bufio"
	"github.com/toffguy77/statusPage/internal/common"
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"os"
	"strings"
)

type SMSProvider struct {
	Name string
}

func (p SMSProvider) GetStatus(countries map[string]models.Country) ([]models.SMSData, error) {
	data, err := parseSmsData(config.SourceData.SMSDataFile)
	if err != nil {
		log.Printf("can't parse prepare data: %v\n", err)
		return nil, err
	}
	result := common.Validate(data, countries)
	return result, nil
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
		if !common.IsEmptyData(dataLine) {
			data = append(data, dataLine)
		}
	}
	return data, nil
}

func parseString(line string) models.SMSData {
	parsedLine := strings.Split(line, ";")
	if !common.IsCorrectLine(parsedLine, 4) {
		log.Printf("line is not valid: %s\n", line)
		return models.SMSData{}
	}
	if !common.IsTrustedProvider(parsedLine[3]) {
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
