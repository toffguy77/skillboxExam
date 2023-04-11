package support

import (
	"encoding/json"
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"io"
	"log"
	"net/http"
)

type SupportProvider struct {
	Name string
}

func (p SupportProvider) GetStatus() ([]models.SupportData, error) {
	data, err := getSupportData(config.SourceData.SupportURL)
	if err != nil {
		log.Printf("can't parse mms data from httpServer: %v\n", err)
		return nil, err
	}
	return data, nil
}

func getSupportData(urlSupportServer string) ([]models.SupportData, error) {
	resp, err := http.Get(urlSupportServer)
	if err != nil {
		log.Printf("error requesting %s for support service data: %v\n", urlSupportServer, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("error response from %s httpServer, status code: %v\n", urlSupportServer, resp.StatusCode)
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response from %s: %v\n", urlSupportServer, err)
		return nil, err
	}

	var data []models.SupportData
	if err = json.Unmarshal(content, &data); err != nil {
		log.Printf("error parsing json response from %s: %v\n", urlSupportServer, err)
		return nil, err
	}

	var result []models.SupportData
	for _, res := range data {
		result = append(result, res)
		//TODO: Fix pointers in slice
	}

	return result, err
}
