package app

import (
    "api-freeradius/models"
    "api-freeradius/internal/repository"
)

func CreateUser(user *models.Radcheck) error {
    user.Op = ":="
    for i := range user.Options {
        user.Options[i].Username = user.Username
        user.Options[i].Op = ":="
    }
    return repository.CreateUser(user)
}

func GetAllUsers() ([]models.Radcheck, error) {
    return repository.GetAllUsers()
}

func GetUser(username string) (*models.Radcheck, error) {
    return repository.GetUser(username)
}

func DeleteUser(username string) error {
    return repository.DeleteUser(username)
}

func UpdateUser(username string, updates *models.Radcheck) error {
    return repository.UpdateUser(username, updates)
}

func DeleteUserOption(username, attribute, value string) error {
    return repository.DeleteUserOption(username, attribute, value)
}
