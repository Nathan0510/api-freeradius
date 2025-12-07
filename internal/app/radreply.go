package app

import (
    "api-freeradius/models"
    "api-freeradius/internal/repository"
)

func CreateUserRadreply(username, attribute, password string) error {

    newUserReply := models.Radreply{
        Username:  username,
        Attribute: attribute,
        Op:        ":=",
        Value:     password,
    }

    return repository.CreateRadreply(&newUserReply)
}

func UpdateUserRadreply(username, attribute, password string) error {

    return repository.UpdateRadreply(username,attribute,password)
}
