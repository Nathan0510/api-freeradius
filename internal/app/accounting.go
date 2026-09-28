package app

import (
    "api-freeradius/models"
    "api-freeradius/db"
)


func GetAccounting(username string) ([]models.Radacct, error) {
    var radacct []models.Radacct
    return radacct, db.DB.Where("username = ?", username).Order("acctstarttime DESC").Find(&radacct).Error
}	