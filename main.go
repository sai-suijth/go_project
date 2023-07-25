package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"name"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	//fmt.Println("Test Line")
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album
	//fmt.Println("insdie add")
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for i := 0; i < len(albums); i += 1 {
		if albums[i].ID == id {
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record not found"})
}

func deleteByTitle(c *gin.Context) {
	name := c.Param("name")
	//fmt.Println(title)
	for i := 0; i < len(albums); i += 1 {
		if albums[i].Title == name {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusAccepted, albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Title to delete doesn't exist"})
}

var albums = []album{
	{ID: "1", Title: "abc", Artist: "artist1", Price: 10},
	{ID: "2", Title: "xyz", Artist: "artist2", Price: 20},
}

func main() {
	fmt.Println("Inside main")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/add", addAlbum)
	router.GET("/getAlbumById/:id", getAlbumById)
	router.DELETE("/deleteByTitle/:name", deleteByTitle)
	router.Run("localhost:8081")
}
