package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	ID        string  `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Location  string  `json:"location"`
	Grade     float64 `json:"grade"`
}

var students = []student{
	{ID: "1", Firstname: "Natsinee", Lastname: "Tana", Location: "Bangkok", Grade: 3.40},
	{ID: "2", Firstname: "Jame", Lastname: "William", Location: "London", Grade: 4.00},
	{ID: "3", Firstname: "Lu", Lastname: "Wang", Location: "China", Grade: 3.10},
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)

	router.Run("localhost:8080")
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}
