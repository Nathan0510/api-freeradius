package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
)

func User(c *gin.Context){

  var json models.User

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := app.CreateUserRadcheck(json.Username, json.Password); err != nil {
    c.JSON(http.StatusOK, gin.H{
      "message": "Erreur ajout radcheck user",
    })
  return
  }

  for _, o := range json.Options {
    if err := app.CreateUserRadreply(json.Username, o.Attribute, o.Value); err != nil {
      c.JSON(http.StatusOK, gin.H{
        "message": "Erreur ajout radreply user",
      })
    return
    }
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Utilisateur"  + json.Username + " bien ajouté",
  })

}
