package migrate

import (
	"github.com/mospyx/data_accounting/pkg/database"
	"github.com/mospyx/data_accounting/pkg/models"
)

func AutoMigrate() error {
	err := database.DB.AutoMigrate(
		&models.Company{},
		&models.CompanyProfile{},
		&models.Employee{},
	)
	if err != nil {
		return err
	}
	return nil
}
