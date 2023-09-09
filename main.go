package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)



const (
	host     = "db"
	port     = 5432
	user     = "admin"
	password = "admin123"
	dbname   = "postgres"
)


var db *sql.DB

type album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {

    initDB()

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/test", sayHiFromAPI)

	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbmus)

	router.PUT("/albums/:id", updateAlbumId)
	router.DELETE("albums/:id", deleteAlbumByID)

	//router.Run("localhost:8080")
	router.Run(":8080")
}

func initDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	createAlbumsTable := `
        CREATE TABLE IF NOT EXISTS albums (
            id SERIAL PRIMARY KEY,
            title VARCHAR(50) NOT NULL,
            artist VARCHAR(50) NOT NULL,
            price float(4) NOT NULL
        )
        `

	_, err = db.Exec(createAlbumsTable)

	if err != nil {
		panic(err)
	}

	fmt.Println("Album tables were created Succesfully")

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Succesfully connected")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbmus(c *gin.Context) {

	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertSQL := "INSERT INTO albums (title, aritst, price) VALUES ($1, $2, $3)"

	_, err := db.Exec(insertSQL, newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})

}

func getAlbumById(c *gin.Context) {

	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album no found"})
}

func updateAlbumId(c *gin.Context) {

	id := c.Param("id")
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {

			a = newAlbum

			albums[i] = newAlbum

			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album was not found"})
}

func deleteAlbumByID(c *gin.Context) {

	id := c.Param("id")

	indexToDelete := -1

	for i, a := range albums {
		if a.ID == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		albums = append(albums[:indexToDelete], albums[indexToDelete+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func sayHiFromAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hi from deez")
}
