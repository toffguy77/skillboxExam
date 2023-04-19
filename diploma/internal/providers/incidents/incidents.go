package incidents

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"io"
	"log"
	"net/http"
)

type IncidentProvider struct {
	Name string
}

func (p IncidentProvider) GetStatus() ([]models.IncidentData, error) {
	data, err := getIncidentList(config.Conf.SourceData.IncidentURL)
	if err != nil {
		log.Printf("can't parse incident data from httpServer: %v\n", err)
		return nil, err
	}
	return data, nil
}

func getIncidentList(urlIncidentServer string) ([]models.IncidentData, error) {
	resp, err := http.Get(urlIncidentServer)
	if err != nil {
		log.Printf("error requesting %s for mms data: %v\n", urlIncidentServer, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("error response from %s httpServer, status code: %v\n", urlIncidentServer, resp.StatusCode)
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response from %s: %v\n", urlIncidentServer, err)
		return nil, err
	}

	var data []models.IncidentData
	if err = json.Unmarshal(content, &data); err != nil {
		log.Printf("error parsing json response from %s: %v\n", urlIncidentServer, err)
		return nil, err
	}

	var result []models.IncidentData
	for _, res := range data {
		if res.Status != "active" && res.Status != "closed" {
			return nil, errors.New(fmt.Sprintf("incident has incorrect status %s: %v\n", res.Status, res))
		}
		result = append(result, res)
	}

	return result, nil
}
