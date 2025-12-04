package app

import (
    "api-freeradius/internal/repository"
)

func UpdateUserRadreply(username, attribute, password string) error {

    return repository.UpdateRadreply(username,attribute,password)
}
