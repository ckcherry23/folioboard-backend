package users

import (
	"github.com/ckcherry23/folioboard-backend/internal/database"
	"github.com/ckcherry23/folioboard-backend/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "Charisma",
		},
	}
	return users, nil
}
