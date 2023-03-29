package countries

import (
	"bufio"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"os"
	"strings"
)

const dataFile = "internal/countries/countries.csv"

func GetCountries() map[string]models.Country {
	countries, err := createCountryMap(dataFile)
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
		if !isEmptryCountry(country) {
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

func isEmptryCountry(country models.Country) bool {
	if country == (models.Country{}) {
		return true
	}
	return false
}
