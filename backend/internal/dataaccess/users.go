package users

import (
	"github.com/CVWO/sample-go-app/backend/internal/database"
	"github.com/CVWO/sample-go-app/backend/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}
