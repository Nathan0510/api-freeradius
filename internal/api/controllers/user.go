// Fichier : user.go

package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
)

// @Summary Create a new Radius user
// @Description Create a new Radius user with radcheck and associated radreply options.
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.Radcheck true "New Radius user data"
// @Success 200 "User added successfully"
// @Failure 400 "Erreur ajout user"
// @Router /users [post]
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

  c.JSON(http.StatusOK, gin.H{"message": "User "  + json.Username + " added successfully"})

}

// @Summary Récupérer tous les utilisateurs Radius
// @Description Récupère la liste complète des utilisateurs (radcheck) avec leurs radreply associées.
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 "Liste des utilisateurs complète"
// @Failure 500 "Erreur interne du serveur"
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
  users, err := app.GetAllUsers()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, users)
}

// @Summary Récupérer un utilisateur Radius spécifique
// @Description Récupère les détails d'un utilisateur spécifique (radcheck) avec ses radreply associées.
// @Tags users 
// @Accept  json
// @Produce  json
// @Param username path string true "Nom d'utilisateur du Radius"
// @Success 200 "Détails de l'utilisateur"
// @Failure 500 "Erreur interne du serveur"
// @Router /users/{username} [get]
func GetUser(c *gin.Context) {
  username := c.Param("username")
  user, err := app.GetUser(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, user)
}

// @Summary Supprimer un utilisateur Radius
// @Description Supprime un utilisateur spécifique (radcheck) et ses radreply associées.
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Nom d'utilisateur du Radius à supprimer"
// @Success 200 "Utilisateur supprimé avec succès"
// @Failure 500 "Erreur interne du serveur"
// @Router /users/{username} [delete]
func DeleteUser(c *gin.Context) {
  username := c.Param("username")
  err := app.DeleteUser(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "User " + username + " deleted successfully"})
}
