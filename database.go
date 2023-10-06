package database

import (
	"uts/Uts-Sds/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root@tcp(localhost:3306)/adminmanagementuser"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})
}

func CreateUser(user *models.User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUser(user *models.User) error {
	result := DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := DB.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func DeleteUserByID(userID uint) error {
	result := DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
