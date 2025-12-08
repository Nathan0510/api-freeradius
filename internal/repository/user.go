package repository

import (
    "gorm.io/gorm"
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateRadcheck(entry *models.Radcheck) error {
    return db.DB.Create(entry).Error
}

func CreateRadreply(entry *models.Radreply) error {
    return db.DB.Create(entry).Error
}

func CreateFullUser(entry *models.Radcheck) error {
    return db.DB.Create(entry).Error
}

func UpdateRadreply(username, attribute, value string) error {
    result := db.DB.
        Model(&models.Radreply{}).
        Where(&models.Radreply{Username: username, Attribute: attribute}).
        Update("Value", value)

    return result.Error
}

func GetAllUsers() ([]models.Radcheck, error) {
    var users []models.Radcheck
    result := db.DB.Preload("Options").Find(&users)
    return users, result.Error
}

func GetUser(username string) (*models.Radcheck, error) {
    var user models.Radcheck
    result := db.DB.Preload("Options").Where("username = ?", username).First(&user)
    return &user, result.Error
}

func DeleteUser(username string) error {
    return db.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Where("username = ?", username).Delete(&models.Radcheck{}).Error; err != nil {
            return err
        }
        return tx.Where("username = ?", username).Delete(&models.Radreply{}).Error
    })
}
