package app

import (
    "api-freeradius/models"
    "api-freeradius/internal/repository"
)

func CreateNas(nasname, shortname, secret string) error {

    nas := models.Nas{
        Nasname:  nasname,
        Shortname: shortname,
        Secret:     secret,
    }

    return repository.CreateNas(&nas)
}

func GetAllNas() ([]models.Nas, error) {
    return repository.GetAllNas()
}

func GetNas(nasname string) (*models.Nas, error) {
    return repository.GetNas(nasname)
}

func DeleteNas(nasname string) error {
    return repository.DeleteNas(nasname)
}

func UpdateNas(nasname string, updates *models.Nas) error {
    return repository.UpdateNas(nasname, updates.Shortname, updates.Secret)
}