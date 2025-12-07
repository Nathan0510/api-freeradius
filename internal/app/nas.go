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
