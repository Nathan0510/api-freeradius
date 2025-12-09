package repository

import (
    "gorm.io/gorm"
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateUser(entry *models.Radcheck) error {
    return db.DB.Create(entry).Error
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

func UpdateUser(username string, updates *models.Radcheck) error {
    return db.DB.Transaction(func(tx *gorm.DB) error {
        if updates.Attribute != "" && updates.Value != "" {
            if err := tx.Model(&models.Radcheck{}).Where("username = ?", username).Updates(map[string]interface{}{
                "Attribute": updates.Attribute,
                "Value":     updates.Value,
            }).Error; err != nil {
                return err
            }
        }

        if len(updates.Options) > 0 {

            newReplies := make([]models.Radreply, len(updates.Options))
            for i, option := range updates.Options {
                newReplies[i] = option
                newReplies[i].Username = username
                newReplies[i].Op = ":="
            }

            if err := tx.Create(&newReplies).Error; err != nil {
                return err
            }
        }

        return nil
    })
}

func DeleteUserOption(username, attribute, value string) error {
    return db.DB.Where("username = ? AND attribute = ? AND value = ?", username, attribute, value,).Delete(&models.Radreply{}).Error
}
