package billing

import (
	"bufio"
	"errors"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"math"
	"os"
)

const BillingDataFile = "billing.data"

type BillingProvider struct {
	Name string
}

func (p BillingProvider) GetStatus() (models.BillingData, error) {
	data, err := parseBillingData(BillingDataFile)
	if err != nil {
		log.Printf("can't parse billing data: %v\n", err)
		return models.BillingData{}, err
	}
	return data, nil
}

func validate(data string) error {
	if len(data) != 6 {
		errIncorrectMask := errors.New("mask contains incorrect number of bits")
		log.Printf("incorrect file content: %v\n", errIncorrectMask)
		return errIncorrectMask
	}
	for _, bitInMask := range data {
		if bitInMask != 48 && bitInMask != 49 {
			errIncorrectMask := errors.New("mask does not contain bits")
			log.Printf("incorrect file content: %v\n", errIncorrectMask)
			return errIncorrectMask
		}
	}
	return nil
}

func parseBillingData(file string) (models.BillingData, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("error opening %s file: %v\n", file, err)
		return models.BillingData{}, err
	}
	defer f.Close()

	var (
		lineCounter int
		dataLine    string
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineCounter++
		if lineCounter > 1 {
			errMoreLines := errors.New("lines than expected")
			log.Printf("incorrect file content: %v\n", errMoreLines)
			return models.BillingData{}, errMoreLines
		}
		dataLine = scanner.Text()

	}
	if err = validate(dataLine); err != nil {
		log.Printf("can't parse billing data: %v\n", err)
		return models.BillingData{}, err
	}
	result := parseMask(dataLine)
	return result, nil
}

func parseMask(mask string) models.BillingData {
	billingStatus := calcNumFromMask(mask)
	createCustomer := 1
	purchase := 2
	payout := 4
	recurring := 8
	fraudControl := 16
	vheckoutPage := 32
	result := models.BillingData{
		CreateCustomer: billingStatus&createCustomer == createCustomer,
		Purchase:       billingStatus&purchase == purchase,
		Payout:         billingStatus&payout == payout,
		Recurring:      billingStatus&recurring == recurring,
		FraudControl:   billingStatus&fraudControl == fraudControl,
		CheckoutPage:   billingStatus&vheckoutPage == vheckoutPage,
	}
	return result
}

func calcNumFromMask(mask string) int {
	maskBytes := []byte(mask)
	var sum int
	for i := 5; i >= 0; i-- {
		if maskBytes[i] == 49 {
			sum += int(math.Pow(float64(2), float64(5-i)))
		}
	}
	return sum
}
