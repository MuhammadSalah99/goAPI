package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    
)

type album struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()

    router.GET("/albums", getAlbums)
    router.GET("/test", sayHiFromAPI)
    
    router.GET("/albums/:id", getAlbumById)
    router.POST("/albums", postAlbmus)

    router.PUT("/albums/:id", updateAlbumId)
    router.DELETE("albums/:id", deleteAlbumByID)
    
    router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbmus(c *gin.Context){

    var newAlbum album

    if err:= c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusOK, newAlbum)
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

     if err:= c.BindJSON(&newAlbum); err != nil {
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
    c.IndentedJSON(http.StatusOK, "Hi from API")
}

