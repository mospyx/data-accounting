package models

import (
	"github.com/mospyx/data_accounting/pkg/database"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID               uint           `gorm:"primarykey" example:"1"`
	CreatedAt        time.Time      `example:"2021-05-27T12:46:29+03:00"`
	UpdatedAt        time.Time      `example:"2021-05-27T12:46:29+03:00"`
	DeletedAt        gorm.DeletedAt `swaggertype:"string" example:"2021-05-27T12:46:29+03:00"`
	CompanyProfileID uint           `json:"company_profile_id" example:"1"`
	CompanyProfile   CompanyProfile `gorm:"association_autocreate:false;association_autoupdate:false"`
}

func (c *Company) Create() error {
	if err := database.DB.Model(&c).Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Company) Update() error {
	if err := database.DB.Model(&c).Updates(&c).Error; err != nil {
		return err
	}
	return nil
}

func GetCompany(id uint) (Company, error) {
	cmp := Company{}
	cmp.ID = id
	if err := database.DB.Set("gorm:auto_preload", true).Take(&cmp).Error; err != nil {
		return Company{}, err
	}
	return cmp, nil
}
