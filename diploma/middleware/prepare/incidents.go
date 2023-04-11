package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/incidents"
	"log"
)

func GetIncidentsData(c chan []models.IncidentData) {
	defer close(c)
	incidentProvider := incidents.IncidentProvider{
		Name: "Incidents List",
	}
	incidentRes, err := incidentProvider.GetStatus()
	if err != nil {
		log.Printf("can't get the list of Incidents: %v", err)
		c <- nil
	}

	c <- sorted(incidentRes)
}

func sorted(res []models.IncidentData) []models.IncidentData {
	for i := 0; i < len(res)-1; i++ {
		for j := i; j < len(res); j++ {
			if res[j].Status == "active" && res[i].Status != "active" {
				res[j], res[i] = res[i], res[j]
				continue
			}
		}
	}
	return res
}
