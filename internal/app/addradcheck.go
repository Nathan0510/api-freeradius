package app

import (
    "api-freeradius/models"
    "api-freeradius/internal/repository"
)

func CreateUserRadcheck(username, password string) error {

    newUserCheck := models.Radcheck{
        Username:  username,
        Attribute: "Cleartext-Password",
        Op:        ":=",
        Value:     password,
    }

    return repository.CreateRadcheck(&newUserCheck)
}
