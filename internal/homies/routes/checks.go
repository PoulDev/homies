package routes

import (
	"github.com/zibbadies/homies/internal/homies/checks"
	"github.com/zibbadies/homies/internal/homies/models"
	"github.com/gin-gonic/gin"
)

func checksDataset(c *gin.Context) {
	res_checks := make(map[string]models.Check)

	for key, value := range checks.CheckersData {
		res_checks[key] = value.Check
	}

	c.JSON(200, res_checks)
}
