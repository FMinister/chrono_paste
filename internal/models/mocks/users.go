package mocks

import (
	"time"

	"github.com/FMinister/chrono_paste/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "duplicate@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (models.User, error) {
	switch id {
	case 1:
		user := models.User{
			ID:      1,
			Name:    "Alice",
			Email:   "alice@example.com",
			Created: time.Now(),
		}
		return user, nil
	default:
		return models.User{}, models.ErrNoRecord
	}
}