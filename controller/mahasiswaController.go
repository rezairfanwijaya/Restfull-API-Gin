package controller

import (
	"log"
	"net/http"
	"restfull-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// bikin struct sebagai penampung inputan dari user untuk table mahasiswa
// yang mana struct MahasiswaInput ini memiliki atribut yang sama seperti Mahasiswa pada package models
type MahasiswaInput struct {
	Nim     string `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}

// function untuk menampilkan semua data mahasiswa
func TampilData(c *gin.Context) {
	// mengambil koneksi database dari context
	db := c.MustGet("db").(*gorm.DB)

	// inisiasi varible baru sebagai representasi struct Mahasiswa
	var mhs []models.Mahasiswa

	// query ke database
	db.Find(&mhs)

	// kirim response ke user
	c.JSON(http.StatusOK, gin.H{
		"data": mhs,
	})

	// kirim log
	log.Println("Berhasil menampilkan semua data mahasiswa")

}
