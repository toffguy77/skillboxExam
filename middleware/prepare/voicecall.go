package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/voicecall"
	"log"
)

func GetVoicecallData(countriesList map[string]models.Country) []models.VoiceCallData {

	voiceProvider := voicecall.VoiceProvider{
		Name: "Voice Calls",
	}
	voiceRes, err := voiceProvider.GetStatus(countriesList)
	if err != nil {
		log.Printf("can't get VoiceCalls status: %v", err)
		return nil
	}
	return voiceRes
}
