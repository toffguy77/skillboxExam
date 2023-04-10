package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/sms"
	"log"
)

func GetSmsData(c chan [][]models.SMSData, countriesList map[string]models.Country) {
	defer close(c)
	smsProvider := sms.SMSProvider{
		Name: "SMS",
	}

	smsRes, err := smsProvider.GetStatus(countriesList)
	if err != nil {
		log.Printf("can't get SMS status: %v", err)
		c <- nil
	}

	smsRes = changeCountrySMS(smsRes, countriesList)

	var res [][]models.SMSData
	res = append(res, sortSmsByProvider(smsRes), sortSmsByCountry(smsRes))

	c <- res
}

func changeCountrySMS(smsRes []models.SMSData, countriesList map[string]models.Country) []models.SMSData {
	var newSmsRes []models.SMSData
	for _, stat := range smsRes {
		stat.Country = countriesList[stat.Country].Name
		newSmsRes = append(newSmsRes, stat)
	}
	return newSmsRes
}

func sortSmsByCountry(smsRes []models.SMSData) []models.SMSData {
	for i := 0; i < len(smsRes)-1; i++ {
		for j := i + 1; j < len(smsRes); j++ {
			if smsRes[i].Country > smsRes[j].Country {
				smsRes[i], smsRes[j] = smsRes[j], smsRes[i]
			}

		}
	}
	return smsRes
}

func sortSmsByProvider(smsRes []models.SMSData) []models.SMSData {
	for i := 0; i < len(smsRes)-1; i++ {
		for j := i + 1; j < len(smsRes); j++ {
			if smsRes[i].Provider > smsRes[j].Provider {
				smsRes[i], smsRes[j] = smsRes[j], smsRes[i]
			}

		}
	}
	return smsRes
}
