package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/mms"
	"log"
)

func GetMmsData(countriesList map[string]models.Country) [][]models.MMSData {

	mmsProvider := mms.MMSProvider{
		Name: "MMS",
	}
	mmsRes, err := mmsProvider.GetStatus(countriesList)
	if err != nil {
		log.Printf("can't get MMS status: %v", err)
		return nil
	}

	mmsRes = changeCountryMMS(mmsRes, countriesList)

	var res [][]models.MMSData
	res = append(res, sortByProvider(mmsRes), sortByCountry(mmsRes))

	return res
}

func changeCountryMMS(mmsRes []models.MMSData, countriesList map[string]models.Country) []models.MMSData {
	var newMmsRes []models.MMSData
	for _, stat := range mmsRes {
		stat.Country = countriesList[stat.Country].Name
		newMmsRes = append(newMmsRes, stat)
	}
	return newMmsRes
}

func sortByCountry(mmsRes []models.MMSData) []models.MMSData {
	for i := 0; i < len(mmsRes)-1; i++ {
		for j := i + 1; j < len(mmsRes); j++ {
			if mmsRes[i].Country > mmsRes[j].Country {
				mmsRes[i].Country, mmsRes[j].Country = mmsRes[j].Country, mmsRes[i].Country
			}

		}
	}
	return mmsRes
}

func sortByProvider(mmsRes []models.MMSData) []models.MMSData {
	for i := 0; i < len(mmsRes)-1; i++ {
		for j := i + 1; j < len(mmsRes); j++ {
			if mmsRes[i].Provider > mmsRes[j].Provider {
				mmsRes[i].Provider, mmsRes[j].Provider = mmsRes[j].Provider, mmsRes[i].Provider
			}

		}
	}
	return mmsRes
}
