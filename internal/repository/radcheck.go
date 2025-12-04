package repository

import (
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateRadcheck(entry *models.Radcheck) error {
    return db.DB.Create(entry).Error
}
