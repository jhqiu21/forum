package users

import (
	"backend/internal/database"
	"backend/internal/models"
	"database/sql"
)

func List(db *sql.DB) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
	}
	return users, nil
}

func SaveUser(user models.User) error {
	db, _ := database.GetDB()
	stmt, err := db.Prepare("INSERT INTO users (username, password, email) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.ID, user.Password)
	if err != nil {
		return err
	}
	return nil
}
