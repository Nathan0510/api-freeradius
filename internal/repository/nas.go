package repository

import (
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateNas(entry *models.Nas) error {
    return db.DB.Create(entry).Error
}

func GetAllNas() ([]models.Nas, error) {
    var nasList []models.Nas
    result := db.DB.Find(&nasList)
    return nasList, result.Error
}

func GetNas(nasname string) (*models.Nas, error) {
    var nas models.Nas
    result := db.DB.Where("nasname = ?", nasname).First(&nas)
    return &nas, result.Error
}

func DeleteNas(nasname string) error {
    result := db.DB.Where("nasname = ?", nasname).Delete(&models.Nas{})
    return result.Error
}

func UpdateNas(nasname string, shortname string, secret string) error {
    
    updateFields := make(map[string]interface{})
        
    if shortname != "" {
        updateFields["shortname"] = shortname
    }
    if secret != "" {
        updateFields["secret"] = secret
    }
    if len(updateFields) == 0 {
        return nil
    }
    result := db.DB.Model(&models.Nas{}).Where("nasname = ?", nasname).Updates(updateFields)
    
    if result.Error != nil {
        return result.Error
    }
    return nil
}