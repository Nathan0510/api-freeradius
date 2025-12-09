package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
)

// GetAllNas godoc
// @Summary Add NAS
// @Description Add NAS
// @Tags NAS
// @Accept json
// @Produce json
// @Param nas body models.Nas true "New NAS data"
// @Success 200 "NAS added successfully"
// @Failure 500 "Failure insert data into dabasase"
// @Router /nas [post]
func CreateNas(c *gin.Context){

  var json models.Nas

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := app.CreateNas(json.Nasname, json.Shortname, json.Secret); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Nas " + json.Nasname + " not added successfully"})
  return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Nas "  + json.Nasname + " added successfully"})

}
// GetAllNas godoc
// @Summary Get all NAS
// @Description Get all NAS
// @Tags NAS
// @Accept json
// @Produce json
// @Success 200 ""
// @Failure 500 "Error"
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
// @Summary Get a specific NAS
// @Description Get a specific NAS
// @Tags NAS
// @Accept json
// @Produce json
// @Success 200 ""
// @Failure 500 "Error"
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

// DeleteNas godoc
// @Summary Delete NAS
// @Description Delete NAS
// @Tags NAS
// @Param username path string true "Nasname of the NAS to delete"
// @Success 200 "Nas deleted successfully"
// @Failure 500 "Error"
// @Router /nas/{username} [delete]
func DeleteNas(c *gin.Context) {
  nasname := c.Param("username")
  if err := app.DeleteNas(nasname); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Nas " + json.Nasname + " not deleted successfully"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Nas " + nasname + " deleted successfully"})
}


// UpdateNas godoc
// @Summary Update a NAS
// @Description Update a NAS
// @Tags NAS
// @Accept json
// @Produce json
// @Param username path string true "Nasname of the NAS to update"
// @Param nas body models.Nas true "Data to update the NAS"
// @Success 200 "Nas updated successfully"
// @Failure 500 "Error"
// @Router /nas/{username} [patch]
func UpdateNas(c *gin.Context) {
  nasname := c.Param("username")
  var json models.Nas

  if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  if err := app.UpdateNas(nasname, &json) ;err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Nas " + nasname + " not updated successfully"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Nas " + nasname + " updated successfully"})
}
