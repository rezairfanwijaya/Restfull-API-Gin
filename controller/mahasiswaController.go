package controller

import (
	"fmt"
	"log"
	"net/http"
	"restfull-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// bikin struct sebagai penampung inputan dari user untuk table mahasiswa
// yang mana struct MahasiswaInput ini memiliki atribut yang sama seperti Mahasiswa pada package models
type MahasiswaInput struct {
	Nim     string `json:"nim" binding:"required"`
	Nama    string `json:"nama" binding:"required"`
	Jurusan string `json:"jurusan" binding:"required"`
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

// function untuk menambahkan data mahasiswa
func TambahMahasiswa(c *gin.Context) {
	// ambil koneksi database melalui context
	db := c.MustGet("db").(*gorm.DB)

	// bikin variable penampung inputan user
	var input MahasiswaInput

	// ambil data dari user melalui inputan json dan harus melalui inputan json
	err := c.ShouldBindJSON(&input)

	// cek error (disini akan menggunakan custom error menggunakan bantuan validator)
	if err != nil {

		// bikin variable penampung error
		var myerr []string

		for _, e := range err.(validator.ValidationErrors) {
			// buat pesan error
			errorMessage := fmt.Sprintf("Error on filed:%s, condition:%s", e.Field(), e.ActualTag())
			// append error
			myerr = append(myerr, errorMessage)
		}

		// tampilkan error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": myerr,
		})

		// matikan kode
		return

	}

	// jika tidak ada error maka inputkan data ke database
	newMhs := models.Mahasiswa{
		Nim:     input.Nim,
		Nama:    input.Nama,
		Jurusan: input.Jurusan,
	}

	// proses simpan ke database
	db.Create(&newMhs)

	// tampilkan response ke user
	c.JSON(http.StatusOK, gin.H{
		"status": "Data mahasiswa berhasil ditambahkan",
	})

}
