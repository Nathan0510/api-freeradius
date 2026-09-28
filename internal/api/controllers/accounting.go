package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "api-freeradius/internal/app"
)

// GetAccounting godoc
// @Summary Get a specific accounting record
// @Description Get a specific accounting record
// @Tags Accounting
// @Accept json
// @Produce json
// @Param username path string true "Username of the accounting record"
// @Success 200 "Data of the accounting record"
// @Failure 500 "Error"
// @Router /api/accounting/{username} [get]
func GetAccounting(c *gin.Context) {
  username := c.Param("username")
  radacct, err := app.GetAccounting(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, radacct)
}