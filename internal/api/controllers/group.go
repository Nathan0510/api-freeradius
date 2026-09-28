package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/models"
  "api-freeradius/internal/app"
  "fmt"
)

// CreateGroup godoc
// @Summary Add Group
// @Description Add Group
// @Tags Group
// @Accept json
// @Produce json
// @Param group body models.Radgroupreply true "Group data"
// @Success 200 "Group added successfully"
// @Failure 500 "Error Group not added successfully"
// @Router /api/groups [post]
func CreateGroup(c *gin.Context){

  var json models.Radgroupreply

  if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if err := app.CreateGroup(json.Groupname, json.Attribute, json.Value); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Group " + json.Groupname + " not added successfully"})
  return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Group"  + json.Groupname + " added successfully"})

}
// GetAllGroup godoc
// @Summary Get all Group
// @Description Get all Group
// @Tags Group
// @Accept json
// @Produce json
// @Success 200 "List of all Group"
// @Failure 500 "Error"
// @Router /api/groups [get]
func GetAllGroup(c *gin.Context) {
  groupList, err := app.GetAllGroup()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, groupList)
}


// GetGroup godoc
// @Summary Get a specific Group
// @Description Get a specific Group
// @Tags Group
// @Accept json
// @Produce json
// @Param groupname path string true "Groupname of the Group"
// @Success 200 "Data of the Group"
// @Failure 500 "Error"
// @Router /api/groups/{groupname} [get]
func GetGroup(c *gin.Context) {
  groupname := c.Param("groupname")
  group, err := app.GetGroup(groupname)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, group)
}

// DeleteGroup godoc
// @Summary Delete Group
// @Description Delete Group
// @Tags Group
// @Param groupname path string true "Groupname of the group to delete"
// @Success 200 "Group deleted successfully"
// @Failure 500 "Error"
// @Router /api/groups/{groupname} [delete]
func DeleteGroup(c *gin.Context) {
  groupname := c.Param("groupname")
  fmt.Println(groupname)
  if err := app.DeleteGroup(groupname); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Group " + groupname + " not deleted successfully"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Group " + groupname + " deleted successfully"})
}


// UpdateGroup godoc
// @Summary Update a Group
// @Description Update a Group
// @Tags Group
// @Accept json
// @Produce json
// @Param groupname path string true "Groupname of the Group to update"
// @Param group body models.Radgroupreply true "Data to update the Group"
// @Success 200 "Group updated successfully"
// @Failure 500 "Error"
// @Router /api/groups/{groupname} [patch]
func UpdateGroup(c *gin.Context) {
  groupname := c.Param("groupname")
  var json models.Radgroupreply

  if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  if err := app.UpdateGroup(groupname, &json) ;err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Group " + groupname + " not updated successfully"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Group" + groupname + " updated successfully"})
}
