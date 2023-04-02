package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/voicecall"
	"log"
)

func GetVoiceCallData(c chan []models.VoiceCallData, countriesList map[string]models.Country) {

	voiceProvider := voicecall.VoiceProvider{
		Name: "Voice Calls",
	}
	voiceRes, err := voiceProvider.GetStatus(countriesList)
	if err != nil {
		log.Printf("can't get VoiceCalls status: %v", err)
		c <- nil
	}
	c <- voiceRes
}
