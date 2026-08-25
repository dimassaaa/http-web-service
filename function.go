package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums merespon dengan list dari semua album sebagai JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums menambahkan sebuah album dari JSON yang diterima dari request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Memanggil BindJSON to mengaitkan JSON yang diterima ke
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// menambahkan album baru ke slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID menemukan album dengan ID value yang sama dengan id
// parameter dikirimkan oleh client, kemudian mengembalikan albumnya sebagai response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop menyeluruh seluruh list album, melihat untuk
	// sebuah album dengan ID value yang cocok dengan parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
