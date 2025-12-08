// Fichier : user.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
)

func CreateUser(c *gin.Context){

  var json models.Radcheck

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


    if err := app.CreateFullUser(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Erreur ajout user",
            "error":   err.Error(),
        })
        return
    }

  c.JSON(http.StatusOK, gin.H{
    "message": "Utilisateur"  + json.Username + " bien ajouté",
  })

}

// @Summary Récupérer tous les utilisateurs Radius
// @Description Récupère la liste complète des utilisateurs (radcheck) avec leurs radreply associées.
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Radcheck "Liste des utilisateurs complète"
// @Failure 500 {object} gin.H "Erreur interne du serveur"
// @Router /users [get]

func GetAllUsers(c *gin.Context) {
  users, err := app.GetAllUsers()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
  username := c.Param("username")
  user, err := app.GetUser(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
  username := c.Param("username")
  err := app.DeleteUser(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
