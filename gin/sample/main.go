package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//album struct
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//dummy data
var albums = []album{
	{ID: "1", Title: "golag", Artist: "vicky", Price: 1000},
	{ID: "2", Title: "python", Artist: "mani", Price: 1500.55},
	{ID: "3", Title: "aws", Artist: "diva", Price: 2222},
}

//get func
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//post func
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//get albums based on ID
func getAlbumByID(c *gin.Context) {
	
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

//main func
func main() {
	//gin router
	router := gin.Default()
	//routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
