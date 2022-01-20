package models

import (
	"errors"
	"github.com/mospyx/data-accounting/pkg/database"
	"gorm.io/gorm"
	"log"
	"time"
)

type Employee struct {
	ID              uint           `gorm:"primarykey" example:"1"`
	CreatedAt       time.Time      `example:"2021-05-27T12:46:29+03:00"`
	UpdatedAt       time.Time      `example:"2021-05-27T12:46:29+03:00"`
	DeletedAt       gorm.DeletedAt `swaggertype:"string" example:"2021-05-27T12:46:29+03:00"`
	FamilyName      string         `json:"family_name"`
	GivenName       string         `json:"given_name"`
	Patronymic      string         `json:"patronymic"`
	CompanyID       uint           `json:"company_id"`
	CompanyPosition string         `json:"company_position"`
	Email           string         `json:"email"`
	Phone           string         `json:"phone"`
	Active          bool           `json:"active"`
}

func (e *Employee) Create(id uint) error {
	e.CompanyID = id
	if err := database.DB.Model(&e).Create(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e *Employee) Update() error {
	if err := database.DB.Model(&e).Updates(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e *Employee) Delete() error {
	if err := database.DB.Model(&e).Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func GetEmployee(id uint) (Employee, error) {
	e := Employee{}
	e.ID = id
	if err := database.DB.Set("gorm:auto_preload", true).Take(&e).Error; err != nil {
		return Employee{}, err
	}
	return e, nil
}

func DeactivateEmployee(id uint) error {
	e, err := GetEmployee(id)
	if err != nil {
		log.Println(err)
	}
	if e.ID == 0 {
		return errors.New("employee not found")
	}
	e.Active = false
	if err = database.DB.Save(&e).Error; err != nil {
		return err
	}
	return nil
}

func GetEmployeeListByCompanyID(cmpID uint) ([]Employee, error) {
	eList := []Employee{}
	if err := database.DB.Model(&Employee{}).Where("company_id = ?", cmpID).Find(&eList).Error; err != nil {
		return nil, err
	}
	return eList, nil
}

func GetEmployeeByIDAndCompanyID(id uint, cmpID uint) (Employee, error) {
	e := Employee{}
	if err := database.DB.Model(&Employee{}).Where("id = ? AND company_id = ?", id, cmpID).Find(&e).Error; err != nil {
		return Employee{}, err
	}
	return e, nil
}
