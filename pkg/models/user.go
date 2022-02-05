package models

import (
	"errors"
	"github.com/mospyx/data-accounting/pkg/database"
	"gorm.io/gorm"
	"log"
	"time"
	"unicode"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	GeneralRole UserRole = "general"
)

type User struct {
	ID         uint           `gorm:"primarykey"`
	CreatedAt  time.Time      `example:"2021-05-27T12:46:29+03:00"`
	UpdatedAt  time.Time      `example:"2021-05-27T12:46:29+03:00"`
	DeletedAt  gorm.DeletedAt `swaggertype:"string" example:"2021-05-27T12:46:29+03:00"`
	FamilyName string         `json:"family_name"`
	GivenName  string         `json:"given_name"`
	Patronymic string         `json:"patronymic"`
	Phone      string         `json:"phone"`
	Email      string         `json:"email" gorm:"unique_index;not null"`
	Password   string         `json:"-"`
	Role       UserRole       `json:"role"`
	Active     bool           `json:"active"`
}

func (r UserRole) IsAdmin() bool {
	return r == AdminRole
}

func (r UserRole) IsGeneral() bool {
	return r == GeneralRole
}

func (u *User) Create() error {
	if err := database.DB.Create(u).Error; err != nil {

		return err
	}
	return nil
}

func GetUser(id uint) (User, error) {
	u := User{}
	u.ID = id
	if err := database.DB.Set("gorm:auto_preload", true).Take(&u).Error; err != nil {
		return User{}, err
	}
	return u, nil
}

func ActivateUser(id uint) error {
	u, err := GetUser(id)
	if err != nil {
		log.Println(err)
	}
	if u.ID == 0 {
		return errors.New("user not found")
	}
	u.Active = true
	if err := database.DB.Save(&u).Error; err != nil {
		return err
	}
	log.Printf("User activating has been successful for user with id: %d \n", u.ID)
	return nil
}

func UpdatePassword(id uint, hashPassword string) error {
	u, _ := GetUser(id)
	if u.ID == 0 {
		return errors.New("user not found")
	}
	u.Password = hashPassword
	if err := database.DB.Save(&u).Error; err != nil {
		return err
	}
	log.Printf("Password updating has been successful for user with id: %d \n", u.ID)
	return nil
}

func Login(email string) (*User, error) {
	userDict := User{}
	if err := database.DB.Where("email = ?", email).First(&userDict).Error; err != nil {
		return nil, err
	}
	return &userDict, nil
}

func CheckEmail(email string) bool {
	tmpUsr := User{}
	database.DB.Where("email = ?", email).Take(&tmpUsr)
	if tmpUsr.ID == 0 {
		return true
	}
	return false
}

func CheckUserName(userName string) bool {
	tmpUsr := User{}
	database.DB.Where("user_name = ?", userName).Take(&tmpUsr)
	if tmpUsr.ID == 0 {
		return true
	}
	return false
}

func CheckUserNameAndEmail(userName, email string) bool {
	if CheckUserName(userName) == false && CheckEmail(email) == false {
		return false
	}
	return true
}

func Password(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {

		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}
