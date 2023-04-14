package prepare

import (
	"github.com/toffguy77/statusPage/internal/config"
	"github.com/toffguy77/statusPage/internal/models"
	"github.com/toffguy77/statusPage/internal/providers/support"
	"log"
)

func GetSupportData(c chan []int) {
	defer close(c)
	supportProvider := support.SupportProvider{
		Name: "Support Status",
	}
	supportRes, err := supportProvider.GetStatus()
	if err != nil {
		log.Printf("can't get Support services status: %v", err)
		c <- nil
	}

	_, totalTickets := calcLoadByTopic(supportRes)
	avgLoad := calcLoad(totalTickets)
	queueTime := calcQueueTime(totalTickets)
	c <- []int{avgLoad, queueTime}
}

func calcQueueTime(tickets int) int {
	speed := 60 / config.Conf.SupportVars.CAPACITY
	return speed * tickets
}

func calcLoad(tickets int) int {
	if tickets < config.Conf.SupportVars.THRESHOLD_LOW {
		return 1
	}
	if tickets < config.Conf.SupportVars.THRESHOLD_MEDIUM {
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
