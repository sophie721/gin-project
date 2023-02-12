package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	group := r.Group("albums")
	group.GET("/", getAlbums)
	group.POST("/", createAlbum)
	group.GET("/:id", getAlbum)

	return r
}

func createAlbum(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	if id != "" {
		for _, album := range albums {
			if album.ID == id {
				c.IndentedJSON(http.StatusOK, album)
				return
			}
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
