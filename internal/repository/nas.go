package repository

import (
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateNas(entry *models.Nas) error {
    return db.DB.Create(entry).Error
}