package models

import (
	"database/sql"
	"time"
)

type Chrono struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type ChronoModel struct {
	DB *sql.DB
}

func (m *ChronoModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *ChronoModel) Get(id int) (Chrono, error) {
	return Chrono{}, nil
}

func (m *ChronoModel) Latest() ([]Chrono, error) {
	return nil, nil
}
