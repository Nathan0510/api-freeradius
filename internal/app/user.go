package app

import (
    "api-freeradius/models"
    "api-freeradius/db"
    "gorm.io/gorm"
)

func CreateUser(user *models.Radcheck) error {
    user.Op = ":="
    for i := range user.Options {
        user.Options[i].Username = user.Username
        user.Options[i].Op = ":="
    }
    return db.DB.Create(user).Error
}

func GetAllUsers() ([]models.Radcheck, error) {
    var users []models.Radcheck
    err := db.DB.Preload("Options").Find(&users).Error
    return users, err
}

func GetUser(username string) (*models.Radcheck, error) {
    var user models.Radcheck
    err := db.DB.Preload("Options").Where("username = ?", username).First(&user).Error
    return &user, err
}

func DeleteUser(username string) error {
    return db.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Where("username = ?", username).Delete(&models.Radcheck{}).Error; err != nil {
            return err
        }
        if err := tx.Where("username = ?", username).Delete(&models.Radreply{}).Error; err != nil {
            return err
        }
        return nil
    })
}

func UpdateUser(username string, updates *models.Radcheck) error {
    return db.DB.Transaction(func(tx *gorm.DB) error {
        // Update Radcheck si nécessaire
        if updates.Attribute != "" && updates.Value != "" {
            if err := tx.Model(&models.Radcheck{}).
                Where("username = ?", username).
                Updates(map[string]interface{}{
                    "Attribute": updates.Attribute,
                    "Value":     updates.Value,
                }).Error; err != nil {
                return err
            }
        }

        // Update ou insert des options une par une
        for _, option := range updates.Options {
            option.Username = username
            option.Op = ":="

            // Essaie de mettre à jour
            res := tx.Model(&models.Radreply{}).
                Where("username = ? AND attribute = ?", username, option.Attribute).
                Updates(map[string]interface{}{"value": option.Value})

            if res.Error != nil {
                return res.Error
            }
        }

        return nil
    })
}
func DeleteUserOption(username, attribute, value string) error {
    return db.DB.Where("username = ? AND attribute = ? AND value = ?", username, attribute, value,).Delete(&models.Radreply{}).Error
}
