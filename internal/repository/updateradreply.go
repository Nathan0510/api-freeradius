package repository

import (
    "api-freeradius/db"
    "api-freeradius/models"
)

func UpdateRadreply(username, attribute, value string) error {
    result := db.DB.
        Model(&models.Radreply{}).
        Where(&models.Radreply{Username: username, Attribute: attribute}).
        Update("Value", value)

    return result.Error
}
