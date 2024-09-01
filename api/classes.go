package api

import (
	"abc/db"
	"abc/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func CreateClasses(c *gin.Context) {
	var class classRequest
	bodyReader, err := c.Request.GetBody()
	if err != nil {
		util.HandelError(c, "Request body is not found", err)
		return
	}
	body, err := io.ReadAll(bodyReader)
	if err != nil {
		util.HandelError(c, "Unable to read the request body", err)
		return
	}
	err = json.Unmarshal(body, &class)
	if err != nil {
		util.HandelError(c, "Unable to parse the request body", err)
		return
	}

	if class.Capacity <= 0 {
		util.HandelError(c, "Capacity cannot be less than zero", err)
		return
	}

	slog.Info("Parsed body is", "class", class)

	var classes []db.ClassInventory
	currentDay := class.StartDate.Time
	for currentDay.Before(class.EndDate.Time) {
		classes = append(classes, db.ClassInventory{
			Name:     class.Name,
			Date:     util.Date{Time: currentDay},
			Capacity: class.Capacity,
		})
		currentDay = currentDay.Add(time.Hour * 24)
	}
	err = db.DB.AddClassInventory(classes)
	if err != nil {
		util.HandelError(c, "Unable to add the classes", err)
	}
	c.Status(http.StatusOK)
}

func GetClasses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.DB.GetClasses())
}

type classRequest struct {
	Name      string    `json:"name"`
	StartDate util.Date `json:"start_date"`
	EndDate   util.Date `json:"end_date"`
	Capacity  int       `json:"capacity"`
}
