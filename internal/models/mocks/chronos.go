package mocks

import (
	"time"

	"github.com/FMinister/chrono_paste/internal/models"
)

var mockChrono = models.Chrono{
	ID:      1,
	Title:   "Test Chrono",
	Content: "Test Chrono Content",
	Created: time.Now(),
	Expires: time.Now().Add(time.Hour),
}

type ChronoModel struct{}

func (m *ChronoModel) Insert(title, content string, expires int) (int, error) {
	return 2, nil
}

func (m *ChronoModel) Get(id int) (models.Chrono, error) {
	switch id {
	case 1:
		return mockChrono, nil
	default:
		return models.Chrono{}, models.ErrNoRecord
	}
}

func (m *ChronoModel) Latest() ([]models.Chrono, error) {
	return []models.Chrono{mockChrono}, nil
}
