package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

// album represents data about a record album.
type album struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type database struct {
	albums []album
	mu     sync.Mutex
}

func initialData() []album {
	return []album{
		{Id: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{Id: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{Id: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

var db database

// albums slice to seed record album data.

func main() {
	db.mu.Lock()
	db.albums = initialData()
	db.mu.Unlock()
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/panic", getPanic)
	router.Run(":8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	db.mu.Lock()
	defer db.mu.Unlock()
	c.IndentedJSON(http.StatusOK, db.albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	db.mu.Lock()
	db.albums = append(db.albums, newAlbum)
	db.mu.Unlock()
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(500, err)
	}

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	db.mu.Lock()
	defer db.mu.Unlock()
	for _, a := range db.albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getPanic(c *gin.Context) {
	panic("I don't know what else to do")

}