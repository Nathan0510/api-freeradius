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
    return nasList, db.DB.Find(&nasList).Error
}

func GetNas(nasname string) (*models.Nas, error) {
    var nas models.Nas
    return &nas, db.DB. Where("nasname = ?", nasname).First(&nas).Error
}

func DeleteNas(nasname string) error {
    return db.DB.Where("nasname = ?", nasname).Delete(&models.Nas{}).Error
}

func UpdateNas(nasname string, updates *models.Nas) error {
    return db.DB.Model(&models.Nas{}).Where("nasname = ?", nasname).Updates(updates).Error
}
