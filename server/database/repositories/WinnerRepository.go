package repositories

import (
	database "github.com/piru72/winners-crud/server/config"
	"github.com/piru72/winners-crud/server/database/models"
)

func GetWinners(season, game, position, teamMember string) ([]models.Winner, error) {
	var winners []models.Winner
	query := database.DB.Model(&models.Winner{})

	// Apply filters
	if season != "" {
		query = query.Where("season = ?", season)
	}
	if game != "" {
		query = query.Where("game = ?", game)
	}
	if position != "" {
		query = query.Where("position = ?", position)
	}
	if teamMember != "" {
		query = query.Where("team_member1 = ? OR team_member2 = ?", teamMember, teamMember)
	}

	// Execute the query
	if err := query.Find(&winners).Error; err != nil {
		return nil, err
	}

	return winners, nil
}
func GetWinnerByID(id string) (*models.Winner, error) {
	var winner models.Winner
	if err := database.DB.Where("id = ?", id).First(&winner).Error; err != nil {
		return nil, err
	}
	return &winner, nil
}

func CreateWinner(winner *models.Winner) error {
	if err := database.DB.Create(winner).Error; err != nil {
		return err
	}
	return nil
}

func UpdateWinner(id string, updatedWinner *models.Winner) (*models.Winner, error) {
	var winner models.Winner
	if err := database.DB.Where("id = ?", id).First(&winner).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Model(&winner).Updates(updatedWinner).Error; err != nil {
		return nil, err
	}

	return &winner, nil
}

func PartialUpdateWinner(id string, updates map[string]interface{}) (*models.Winner, error) {
	var winner models.Winner
	if err := database.DB.Where("id = ?", id).First(&winner).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Model(&winner).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &winner, nil
}

func DeleteWinner(id string) error {
	var winner models.Winner
	if err := database.DB.Where("id = ?", id).First(&winner).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&winner).Error; err != nil {
		return err
	}

	return nil
}
