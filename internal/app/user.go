package app

import (
    "api-freeradius/models"
    "api-freeradius/internal/repository"
)

func CreateFullUser(user *models.Radcheck) error {
    user.Op = ":="
    for i := range user.Options {
        user.Options[i].Username = user.Username
        user.Options[i].Op = ":="
    }
    return repository.CreateFullUser(user)
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

func UpdateUserRadreply(username, attribute, password string) error {

    return repository.UpdateRadreply(username,attribute,password)
}
