package email

import (
	"bufio"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"os"
	"strconv"
	"strings"
)

const EmailDataFile = "email.data"

type EmailProvider struct {
	Name string
}

func (p EmailProvider) GetStatus(countries map[string]models.Country) ([]models.EmailData, error) {
	data, err := parseEmailData(EmailDataFile)
	if err != nil {
		log.Printf("can't parse email data: %v\n", err)
		return nil, err
	}
	result := validate(data, countries)
	return result, nil
}

func validate(data []models.EmailData, country map[string]models.Country) []models.EmailData {
	for iter, res := range data {
		_, ok := country[res.Country]
		if !ok {
			log.Printf("email data is not valid: %v\n", res)
			data[iter] = data[len(data)-1]
			data = data[:len(data)-1]
		}
	}
	return data
}

func parseEmailData(file string) ([]models.EmailData, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("error opening %s file: %v\n", file, err)
		return nil, err
	}
	defer f.Close()

	var data []models.EmailData
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dataLine := parseString(scanner.Text())
		if !isEmptyEmail(dataLine) {
			data = append(data, dataLine)
		}
	}
	return data, nil
}

func isEmptyEmail(emailData models.EmailData) bool {
	if emailData == (models.EmailData{}) {
		return true
	}
	return false
}

func parseString(line string) models.EmailData {
	parsedLine := strings.Split(line, ";")
	if !isCorrectLine(parsedLine, 3) {
		log.Printf("line is not valid: %s\n", line)
		return models.EmailData{}
	}
	if !isTrustedemailProvider(parsedLine[1]) {
		log.Printf("untrusted provider, skip: %s\n", line)
		return models.EmailData{}
	}

	deliveryTime, err := strconv.Atoi(parsedLine[2])
	if err != nil {
		log.Printf("error converting DeliveryTime %s data: %v\n", parsedLine[2], err)
		return models.EmailData{}
	}
	data := models.EmailData{
		Country:      parsedLine[0],
		Provider:     parsedLine[1],
		DeliveryTime: deliveryTime,
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

func isTrustedemailProvider(provider string) bool {
	switch strings.ToLower(provider) {
	case
		"gmail", "yahoo", "hotmail", "msn", "orange", "comcast",
		"aol", "live", "rediffmail", "gmx", "protonmail", "yandex", "mail.ru":
		return true
	}
	log.Printf("email provider is not trusted: %s\n", provider)
	return false
}
