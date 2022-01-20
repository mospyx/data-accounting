package models

import (
	"github.com/mospyx/data-accounting/pkg/database"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	ID               uint           `gorm:"primarykey" example:"1"`
	CreatedAt        time.Time      `example:"2021-05-27T12:46:29+03:00"`
	UpdatedAt        time.Time      `example:"2021-05-27T12:46:29+03:00"`
	DeletedAt        gorm.DeletedAt `swaggertype:"string" example:"2021-05-27T12:46:29+03:00"`
	CompanyProfileID uint           `json:"company_profile_id" example:"1"`
	CompanyProfile   CompanyProfile `json:"company_profile"`
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

func (c *Company) Delete() error {
	if err := database.DB.Model(&c).Delete(&c).Error; err != nil {
		return err
	}
	return nil
}

func GetCompany(id uint) (Company, error) {
	cmp := Company{}
	if err := database.DB.Where("id = ?", id).Preload("CompanyProfile").Take(&cmp).Error; err != nil {
		return Company{}, err
	}
	return cmp, nil
}

func GetCompanyList() ([]Company, error) {
	var cmpList []Company
	if err := database.DB.Model(&Company{}).Preload("CompanyProfile").Find(&cmpList).Error; err != nil {
		return nil, err
	}
	return cmpList, nil
}
