package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/refill_events", getRefillEvents)
	router.GET("/refill_events/:id", getRefillEventByID)
	router.POST("/refill_events", postRefillEvent)

	router.GET("/user_events/:user_id", getUserRefillEvents)
	router.GET("/user_events_sum/:user_id", getUserConsumedWater)

	router.Run("localhost:8080")
}

type refill_event struct {
	ID         string    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	Milliliter int       `json:"milliliter"`
	UserID     string    `json:"user_id"`
}

var refill_events = []refill_event{
	{ID: "1", Timestamp: time.Now(), Milliliter: 500, UserID: "1"},
	{ID: "2", Timestamp: time.Now().Add(time.Hour * 24 * 10), Milliliter: 1000, UserID: "1"},
	{ID: "3", Timestamp: time.Date(2024, time.April, 12, 8, 33, 0, 0, time.UTC), Milliliter: 100, UserID: "2"},
}

func getRefillEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, refill_events)
}

func postRefillEvent(c *gin.Context) {
	var newEvent refill_event

	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	refill_events = append(refill_events, newEvent)
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func getRefillEventByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range refill_events {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "refill event not found"})
}

func getUserRefillEvents(c *gin.Context) {
	userId := c.Param("user_id")
	var user_refill_events []refill_event

	for _, a := range refill_events {
		if a.UserID == userId {
			user_refill_events = append(user_refill_events, a)
		}
	}

	if len(user_refill_events) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no events for user found"})
	} else {
		c.IndentedJSON(http.StatusOK, user_refill_events)
	}
}

func getUserConsumedWater(c *gin.Context) {
	userId := c.Param("user_id")
	var waterConsumed int

	for _, a := range refill_events {
		if a.UserID == userId {
			waterConsumed += a.Milliliter
		}
	}

	c.IndentedJSON(http.StatusOK, waterConsumed)
}
