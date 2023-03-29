package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/email"
	"log"
)

func GetEmailData(countriesList map[string]models.Country) map[string][][]models.EmailData {
	emailProvider := email.EmailProvider{
		Name: "E-mails",
	}
	emailRes, err := emailProvider.GetStatus(countriesList)
	if err != nil {
		log.Printf("can't get Email status: %v", err)
		return nil
	}
	emailResByCountry := resByCountry(emailRes)
	emailSortedByDeliveryTime := sortedByDeliveryTime(emailResByCountry)

	return emailSortedByDeliveryTime
}

func sortedByDeliveryTime(emailResByCountry map[string][]models.EmailData) map[string][][]models.EmailData {
	sortedResByCountry := make(map[string][][]models.EmailData)
	for country, data := range emailResByCountry {
		sortedResByCountry[country] = [][]models.EmailData{sortByDeliveryAsc(data), sortByDeliveryDesc(data)}
	}
	return sortedResByCountry
}

func sortByDeliveryDesc(data []models.EmailData) []models.EmailData {
	if len(data) == 1 {
		return data
	}
	for i := 0; i < len(data)-1; i++ {
		for j := i; j < len(data); j++ {
			if data[i].DeliveryTime < data[j].DeliveryTime {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	res := make([]models.EmailData, 3)
	copy(res, data[0:3])
	return res
}

func sortByDeliveryAsc(data []models.EmailData) []models.EmailData {
	if len(data) == 1 {
		return data
	}
	for i := 0; i < len(data)-1; i++ {
		for j := i; j < len(data); j++ {
			if data[i].DeliveryTime > data[j].DeliveryTime {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	res := make([]models.EmailData, 3)
	copy(res, data[0:3])
	return res
}

func resByCountry(results []models.EmailData) map[string][]models.EmailData {
	sortedResByCountry := make(map[string][]models.EmailData)
	for _, res := range results {
		if _, ok := sortedResByCountry[res.Country]; ok {
			sortedResByCountry[res.Country] = append(sortedResByCountry[res.Country], res)
		} else {
			sortedResByCountry[res.Country] = []models.EmailData{res}
		}
	}
	return sortedResByCountry
}
