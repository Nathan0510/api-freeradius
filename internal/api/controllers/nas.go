package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
)


func CreateNas(c *gin.Context){

  var json models.Nas

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := app.CreateNas(json.Nasname, json.Shortname, json.Secret); err != nil {
    c.JSON(http.StatusOK, gin.H{
      "message": "Erreur ajout nas user",
    })
  return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Nas"  + json.Nasname + " bien ajouté",
  })

}
// GetAllNas godoc
// @Summary Récupérer tous les NAS Radius
// @Description Récupère la liste complète de tous les NAS configurés
// @Tags NAS
// @Accept json
// @Produce json
// @Success 200 {array} models.Nas "Liste des NAS récupérée avec succès"
// @Failure 500 {object} map[string]interface{} "Erreur serveur interne"
// @Router /nas [get]
func GetAllNas(c *gin.Context) {
  nasList, err := app.GetAllNas()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, nasList)
}


// GetNas godoc
// @Summary Afficher un NAS
// @Description Afficher la configuration d'un NAS
// @Tags NAS
// @Accept json
// @Produce json
// @Success 200 {array} models.Nas "NAS récupérée avec succès"
// @Failure 500 {object} map[string]interface{} "Erreur serveur interne"
// @Router /nas/{username} [get]
func GetNas(c *gin.Context) {
  nasname := c.Param("username")
  nas, err := app.GetNas(nasname)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, nas)
}

// DeleteNas supprime un NAS par son nom d'utilisateur
// @Summary Supprime un NAS
// @Description Supprime un NAS existant à partir du nom d'utilisateur fourni
// @Tags NAS
// @Param username path string true "Nom d'utilisateur du NAS"
// @Success 200 {object} map[string]string "Message de succès"  example(map[string]string{"message":"Nas deleted successfully"})
// @Failure 500 {object} map[string]string "Message d'erreur"
// @Router /nas/{username} [delete]
func DeleteNas(c *gin.Context) {
  nasname := c.Param("username")
  err := app.DeleteNas(nasname)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Nas deleted successfully"})
}

func UpdateNas(c *gin.Context) {
  nasname := c.Param("username")
  var json models.NewNas

  if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  err := app.UpdateNas(nasname, &json)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Nas updated successfully"})
}
