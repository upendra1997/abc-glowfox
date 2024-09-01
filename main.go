package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT int = 8080

type Class struct {
	Name      string `json:"name"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

func HandelError(c *gin.Context, msg string, err error) {
	if err != nil {
		slog.Error(msg, "err", err)
		c.IndentedJSON(http.StatusBadRequest, map[string]string{
			"error": msg,
		})
	}
}

func CreateClasses(c *gin.Context) {
	var class Class
	bodyReader, err := c.Request.GetBody()
	if err != nil {
		HandelError(c, "Request body is not found", err)
	}
	body, err := io.ReadAll(bodyReader)
	if err != nil {
		HandelError(c, "Unable to read the request body", err)
	}
	err = json.Unmarshal(body, &class)
	HandelError(c, "Unable to parse the request body", err)
	slog.Info("Parsed body is", "class", class)
	c.IndentedJSON(http.StatusOK, Class{
		Name:      "Upendra Upadhyay",
		StartDate: 20,
		EndDate:   30,
	})
}

func GetClassses(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		slog.Info("name not found")
	}
	start_date, ok := c.GetQuery("start_date")
	if !ok {
		slog.Info("start_date not found")
	}
	end_date, ok := c.GetQuery("end_date")
	if !ok {
		slog.Info("end_date not found")
	}
	slog.Info("values parsed from the query param", "name", name, "start_date", start_date, "end_date", end_date)
	c.IndentedJSON(http.StatusOK, Class{
		Name:      "Upendra Upadhyay",
		StartDate: 20,
		EndDate:   30,
	})
}

func handleBookings(rw http.ResponseWriter, r *http.Request) {

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes := r.Group("/api")
	{
		routes.GET("/classes", GetClassses)
		routes.POST("/classes", CreateClasses)
	}
	return r
}

func main() {
	r := setupRouter()
	err := r.Run(fmt.Sprintf(":%d", PORT))
	if err != nil {
		slog.Error("Server Crashed", "err", err)
	}
}
