package main

import (
	"FYP_TravelPlanner/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	server := gin.Default()
	server.GET("/plans", getEvents)
	server.POST("/plans", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	itineraries := models.GetAllItineraries()
	context.JSON(http.StatusOK, itineraries)
}

func createEvent(context *gin.Context) {
	var Itinerary models.Itinerary
	err := context.ShouldBindJSON(&Itinerary)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	Itinerary.ID = 1
	Itinerary.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"Message": "Itinerary created"})
}
