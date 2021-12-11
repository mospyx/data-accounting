package migrate

import "github.com/mospyx/data_accounting/pkg/database"

func AutoMigrate() error {
	err := database.DB.AutoMigrate(
	//&struct
	)
	if err != nil {
		return err
	}
	return nil
}
