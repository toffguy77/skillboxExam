package mms

import (
	"encoding/json"
	"github.com/toffguy77/statusPage/internal/models"
	"io"
	"log"
	"net/http"
	"strings"
)

type MMSProvider struct {
	Name string
}

func (p MMSProvider) GetStatus(countries map[string]models.Country) ([]models.MMSData, error) {
	data, err := getMmsData("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Printf("can't parse mms data from httpServer: %v\n", err)
		return nil, err
	}
	result := validate(data, countries)
	return result, nil
}

func validate(data []models.MMSData, country map[string]models.Country) []models.MMSData {
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

func getMmsData(urlMmsServer string) ([]models.MMSData, error) {
	resp, err := http.Get(urlMmsServer)
	if err != nil {
		log.Printf("error requesting %s for mms data: %v\n", urlMmsServer, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("error response from %s httpServer, status code: %v\n", urlMmsServer, resp.StatusCode)
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response from %s: %v\n", urlMmsServer, err)
		return nil, err
	}

	var data []models.MMSData
	if err = json.Unmarshal(content, &data); err != nil {
		log.Printf("error parsing json response from %s: %v\n", urlMmsServer, err)
		return nil, err
	}

	var result []models.MMSData
	for _, res := range data {
		if isTrustedMmsProvider(res.Provider) {
			result = append(result, res)
		}
	}

	return result, err
}

func isTrustedMmsProvider(provider string) bool {
	switch strings.ToLower(provider) {
	case
		"topolo", "rond", "kildy":
		return true
	}
	log.Printf("mms provider is not trusted: %s\n", provider)
	return false
}
