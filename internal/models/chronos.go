package models

import (
	"database/sql"
	"errors"
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
	stmt := `INSERT INTO chronos (title, content, created, expires) 
			VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + '1 DAY'::INTERVAL * $3)
			RETURNING id`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ChronoModel) Get(id int) (Chrono, error) {
	stmt := `SELECT id, title, content, created, expires FROM chronos 
			WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	var c Chrono
	err := m.DB.QueryRow(stmt, id).Scan(&c.ID, &c.Title, &c.Content, &c.Created, &c.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Chrono{}, ErrNoRecord
		} else {
			return Chrono{}, err
		}
	}

	return c, nil
}

func (m *ChronoModel) Latest() ([]Chrono, error) {
	return nil, nil
}
