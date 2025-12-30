package app

import (
    "api-freeradius/models"
    "api-freeradius/db"
)

func CreateNas(nasname, shortname, secret string) error {
    nas := &models.Nas{
        Nasname:  nasname,
        Shortname: shortname,
        Secret:     secret,
    }

    return db.DB.Create(nas).Error
}

func GetAllNas() ([]models.Nas, error) {
    var nasList []models.Nas
    err := db.DB.Find(&nasList).Error
    return nasList, err
}

func GetNas(nasname string) (*models.Nas, error) {
    var nas models.Nas
    err := db.DB. Where("nasname = ?", nasname).First(&nas).Error
    return &nas, err
}

func DeleteNas(nasname string) error {
    return db.DB.Where("nasname = ?", nasname).Delete(&models.Nas{}).Error
}

func UpdateNas(nasname string, updates *models.Nas) error {
    result := db.DB.Model(&models.Nas{}).Where("nasname = ?", nasname).Updates(updates)

    if result.Error != nil {
        return result.Error
    }

    return nil
}