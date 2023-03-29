package providers

import "github.com/toffguy77/statusPage/internal/models"

type Provider interface {
	GetStatus(map[string]*models.Country) string
}
