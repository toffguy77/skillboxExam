package countries

import (
	"bufio"
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"os"
	"strings"
)

func GetCountries() map[string]models.Country {
	countries, err := createCountryMap(config.Conf.SourceData.CountriesDataFile)
	if err != nil {
		log.Fatal("can't load country list file\n")
	}
	return countries
}

func createCountryMap(file string) (map[string]models.Country, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("error opening %s file: %v", file, err)
		return nil, err
	}
	defer f.Close()

	data := make(map[string]models.Country)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		country := parseString(scanner.Text())
		if !isEmptyCountry(country) {
			data[country.Alpha2] = country
		}
	}
	return data, nil
}

func parseString(line string) models.Country {
	parsedLine := strings.Split(line, ";")
	if len(parsedLine) != 5 || len(parsedLine[1]) != 2 || len(parsedLine[2]) != 3 {
		log.Printf("line is not valid: %s", line)
		return models.Country{}
	}
	country := models.Country{
		Name:       parsedLine[0],
		Alpha2:     parsedLine[1],
		Alpha3:     parsedLine[2],
		ISO_3166_1: parsedLine[3],
		ISO_3166_2: parsedLine[4],
	}
	return country
}

func isEmptyCountry(country models.Country) bool {
	if country.Name == "" || country.Alpha2 == "" || country.Alpha3 == "" || country.ISO_3166_1 == "" || country.ISO_3166_2 == "" {
		return true
	}
	return false
}
