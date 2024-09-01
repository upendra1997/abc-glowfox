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
	bodyReader := c.Request.Body
	body, err := io.ReadAll(bodyReader)
	if err != nil {
		util.HandleError(c, "Unable to read the request body", err)
		return
	}
	err = json.Unmarshal(body, &class)
	if err != nil {
		util.HandleError(c, "Unable to parse the request body", err)
		return
	}

	if class.Capacity <= 0 {
		util.HandleError(c, "Capacity cannot be less than zero", err)
		return
	}

	class.EndDate = util.Date{Time: class.EndDate.Add(time.Hour * 24)}
	slog.Info("Parsed body is", "class", class)

	duration := class.EndDate.Time.Sub(class.StartDate.Time).Hours() / 24

	if duration < 0 {
		util.HandleError(c, "please specify valid start and end date", err)
		return
	}

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
		util.HandleError(c, "Unable to add the classes", err)
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
