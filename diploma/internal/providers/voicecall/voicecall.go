package voicecall

import (
	"bufio"
	"github.com/toffguy77/statusPage/internal/common"
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"log"
	"os"
	"strconv"
	"strings"
)

type VoiceProvider struct {
	Name string
}

func (p VoiceProvider) GetStatus(countries map[string]models.Country) ([]models.VoiceCallData, error) {
	data, err := parseVoiceCallData(config.Conf.SourceData.VoiceCallDataFile)
	if err != nil {
		log.Printf("can't parse voice data: %v\n", err)
		return nil, err
	}
	result := common.Validate(data, countries)
	return result, nil
}

func parseVoiceCallData(file string) ([]models.VoiceCallData, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("error opening %s file: %v\n", file, err)
		return nil, err
	}
	defer f.Close()

	var data []models.VoiceCallData
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dataLine := parseString(scanner.Text())
		if !isEmptyVoiceData(dataLine) {
			data = append(data, dataLine)
		}
	}
	return data, nil
}

func isEmptyVoiceData(data models.VoiceCallData) bool {
	if data == (models.VoiceCallData{}) {
		return true
	}
	return false
}

func parseString(line string) models.VoiceCallData {
	parsedLine := strings.Split(line, ";")
	if !isCorrectLine(parsedLine, 8) {
		log.Printf("line is not valid: %s\n", line)
		return models.VoiceCallData{}
	}
	if !isTrustedVoiceProvider(parsedLine[3]) {
		log.Printf("untrusted provider, skip: %s\n", line)
		return models.VoiceCallData{}
	}
	connectionStability, err := strconv.ParseFloat(parsedLine[4], 32)
	if err != nil {
		log.Printf("error converting ConnectionStability %s data: %v\n", parsedLine[5], err)
		return models.VoiceCallData{}
	}
	ttfb, err := strconv.Atoi(parsedLine[5])
	if err != nil {
		log.Printf("error converting TTFB %s data: %v\n", parsedLine[5], err)
		return models.VoiceCallData{}
	}
	voicePurity, err := strconv.Atoi(parsedLine[6])
	if err != nil {
		log.Printf("error converting VoicePurity %s data: %v\n", parsedLine[6], err)
		return models.VoiceCallData{}
	}
	medianOfCallsTime, err := strconv.Atoi(parsedLine[7])
	if err != nil {
		log.Printf("error converting MedianOfCallsTime %s data: %v\n", parsedLine[7], err)
		return models.VoiceCallData{}
	}

	data := models.VoiceCallData{
		Country:             parsedLine[0],
		Bandwidth:           parsedLine[1],
		ResponseTime:        parsedLine[2],
		Provider:            parsedLine[3],
		ConnectionStability: float32(connectionStability),
		TTFB:                ttfb,
		VoicePurity:         voicePurity,
		MedianOfCallsTime:   medianOfCallsTime,
	}
	return data
}

func isCorrectLine(line []string, l int) bool {
	if len(line) != l {
		return false
	}
	for _, val := range line {
		if val == "" {
			return false
		}
	}
	return true
}

func isTrustedVoiceProvider(provider string) bool {
	switch strings.ToLower(provider) {
	case
		"transparentcalls", "e-voice", "justphone":
		return true
	}
	log.Printf("prepare provider is not trusted: %s\n", provider)
	return false
}
