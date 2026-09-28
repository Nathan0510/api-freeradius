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
// @Failure 500 "Error"
// @Router /api/users [post]
func CreateUser(c *gin.Context){

  var json models.Radcheck

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := app.CreateUser(&json); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "User "  + json.Username + " added successfully"})

}

// @Summary Get all radius users
// @Description Get all radius users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 "List of all user"
// @Failure 500 "Error"
// @Router /api/users [get]
func GetAllUsers(c *gin.Context) {
  users, err := app.GetAllUsers()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, users)
}

// @Summary Get specific radius user
// @Description Get specific radius user
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Username of radius user"
// @Success 200 ""
// @Failure 500 "Error"
// @Router /api/users/{username} [get]
func GetUser(c *gin.Context) {
  username := c.Param("username")
  user, err := app.GetUser(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, user)
}

// @Summary Delete radius user
// @Description Delete radius user
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Username of radius user"
// @Success 200 "User deleted successfully"
// @Failure 500 "Error"
// @Router /api/users/{username} [delete]
func DeleteUser(c *gin.Context) {
  username := c.Param("username")
  err := app.DeleteUser(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "User " + username + " deleted successfully"})
}

// @Summary Update radius user
// @Description Update radius user
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Username of radius user"
// @Param user body models.Radcheck true "Radius user data"
// @Success 200 "User uptaded successfully"
// @Failure 500 "Error"
// @Router /api/users/{username} [patch]
func UpdateUser(c *gin.Context) {
    username := c.Param("username")
    var json models.Radcheck

    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := app.UpdateUser(username, &json); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User " + username + " updated successfully"})
}

// @Summary Delete option of radius user
// @Description Delete option of radius user
// @Tags users
// @Accept  json
// @Produce  json
// @Param username path string true "Username of radius user"
// @Param option body models.Radreply true "Radius user option data"
// @Success 200 "Option user deleted successfully"
// @Failure 500 "Error"
// @Router /api/users/{username}/options/{optionname} [delete]
func DeleteUserOption(c *gin.Context) {
  username := c.Param("username")
  attribute := c.Param("optionname")
  value := c.Query("value")

  if value == "" {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Value is required as query parameter"})
      return
  }
  if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := app.DeleteUserOption(username, attribute, value); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
    c.JSON(http.StatusOK, gin.H{"message": "Option " + attribute +  " value " + value + " user " + username + " deleted successfully"})
}
