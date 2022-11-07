package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db *sql.DB
var albums []album

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/vinylshop?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	//	check connection
	if err = db.Ping(); err != nil {
		log.Println(err)
	} else {
		fmt.Println("successfully connected to database")
	}
}

func main() {
	router := gin.Default()
	http.Handle("/favicon", http.NotFoundHandler())

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	http.ListenAndServe(":8080", router)
}

func getAlbums(c *gin.Context) {
	albums = []album{}
	rows, err := db.Query("select * from albums")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	//	iterate through results
	for rows.Next() {
		alb := album{}
		err = rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			log.Println(err)
		}
		albums = append(albums, alb)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	err := c.MustBindWith(&newAlbum, binding.JSON)
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("insert into albums values (default, $1, $2, $3)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		log.Println(err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	alb := album{}
	// query the db by ID
	row := db.QueryRow("select * from albums where ID = $1", idInt)
	err = row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, alb)
}
