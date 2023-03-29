package prepare

import (
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/support"
	"log"
)

var (
	THRESHOLD_LOW    = 9
	THRESHOLD_MEDIUM = 16
	CAPACITY         = 18
	SPECIALISTS      = 7
)

func GetSupportData() []int {
	supportProvider := support.SupportProvider{
		Name: "Support Status",
	}
	supportRes, err := supportProvider.GetStatus()
	if err != nil {
		log.Printf("can't get Support services status: %v", err)
		return nil
	}

	_, totalTickets := calcLoadByTopic(supportRes)
	avgLoad := calcLoad(totalTickets)
	queueTime := calcQueueTime(totalTickets)
	return []int{avgLoad, queueTime}
}

func calcQueueTime(tickets int) int {
	speed := 60 / CAPACITY
	return speed * tickets
}

func calcLoad(tickets int) int {
	if tickets < THRESHOLD_LOW {
		return 1
	}
	if tickets < THRESHOLD_MEDIUM {
		return 2
	}
	return 3
}

func calcLoadByTopic(supportRes []models.SupportData) (map[string]int, int) {
	loadByTopic := make(map[string]int)
	totalTickets := 0
	for _, data := range supportRes {
		if _, ok := loadByTopic[data.Topic]; ok {
			loadByTopic[data.Topic] += data.ActiveTickets
		} else {
			loadByTopic[data.Topic] = data.ActiveTickets
		}
		totalTickets += data.ActiveTickets
	}
	return loadByTopic, totalTickets
}
