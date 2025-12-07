package repository

import (
    "api-freeradius/db"
    "api-freeradius/models"
)

func CreateRadreply(entry *models.Radreply) error {
    return db.DB.Create(entry).Error
}

func UpdateRadreply(username, attribute, value string) error {
    result := db.DB.
        Model(&models.Radreply{}).
        Where(&models.Radreply{Username: username, Attribute: attribute}).
        Update("Value", value)

    return result.Error
}
