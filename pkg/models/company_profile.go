package models

import (
	"github.com/mospyx/data-accounting/pkg/database"
	"gorm.io/gorm"
	"time"
)

type CompanyProfile struct {
	ID                uint           `gorm:"primarykey" example:"1"`
	CreatedAt         time.Time      `example:"2021-05-27T12:46:29+03:00"`
	UpdatedAt         time.Time      `example:"2021-05-27T12:46:29+03:00"`
	DeletedAt         gorm.DeletedAt `swaggertype:"string" example:"2021-05-27T12:46:29+03:00"`
	Name              string         `json:"name"`
	NumberOfEmployees string         `json:"number_of_employees"`
	//Employees         []Employee     `json:"employees"`
	Website string `json:"website"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func (cp *CompanyProfile) Create() error {
	if err := database.DB.Model(&cp).Create(&cp).Error; err != nil {
		return err
	}
	return nil
}

func (cp *CompanyProfile) Update() error {
	if err := database.DB.Model(&cp).Updates(&cp).Error; err != nil {
		return err
	}
	return nil
}

func (cp *CompanyProfile) Delete() error {
	if err := database.DB.Model(&cp).Delete(&cp).Error; err != nil {
		return err
	}
	return nil
}
