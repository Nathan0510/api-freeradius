package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
)

func Nas(c *gin.Context){

  var json models.Nas

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := app.CreateNas(json.Nasname, json.Shortname, json.Secret); err != nil {
    c.JSON(http.StatusOK, gin.H{
      "message": "Erreur ajout radcheck user",
    })
  return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Nas"  + json.Username + " bien ajouté",
  })

}
