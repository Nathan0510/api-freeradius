package repository

import (
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateRadreply(entry *models.Radreply) error {
    return db.DB.Create(entry).Error
}
