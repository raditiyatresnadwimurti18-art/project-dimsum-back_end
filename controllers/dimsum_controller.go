package controllers

import (
	"context"
	"net/http"

	// Sesuaikan "project-dimsum" dengan nama module di file go.mod kamu
	"project-dimsum/config"
	"project-dimsum/models"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// Nama fungsi diawali huruf kapital
func TambahDimsum(c *gin.Context) {
	var input models.Dimsum

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	ctx := context.Background()
	ref, _, err := config.DB.Collection("dimsums").Add(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan ke database"})
		return
	}

	input.ID = ref.ID
	c.JSON(http.StatusCreated, input)
}

func GetSemuaDimsum(c *gin.Context) {
	ctx := context.Background()
	var daftarDimsum []models.Dimsum

	iter := config.DB.Collection("dimsums").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
			return
		}

		var d models.Dimsum
		doc.DataTo(&d)
		d.ID = doc.Ref.ID

		daftarDimsum = append(daftarDimsum, d)
	}

	c.JSON(http.StatusOK, daftarDimsum)
}

func UpdateDimsum(c *gin.Context) {
	id := c.Param("id")
	var input models.Dimsum

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	ctx := context.Background()
	_, err := config.DB.Collection("dimsums").Doc(id).Set(ctx, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data"})
		return
	}

	input.ID = id
	c.JSON(http.StatusOK, input)
}

func HapusDimsum(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()
	_, err := config.DB.Collection("dimsums").Doc(id).Delete(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data dimsum berhasil dihapus"})
}
