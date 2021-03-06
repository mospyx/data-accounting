package migrate

import (
	"github.com/mospyx/data-accounting/pkg/database"
	"github.com/mospyx/data-accounting/pkg/models"
)

func AutoMigrate() error {
	err := database.DB.AutoMigrate(
		&models.Company{},
		&models.CompanyProfile{},
		&models.Employee{},
		&models.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
