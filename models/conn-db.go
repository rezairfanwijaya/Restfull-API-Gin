package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/restfull_api_gin_go?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error
	if err != nil {
		log.Panicln("Gagal terhubung ke database")
		panic(err)
	}

	// jika berhasil maka lakukan migration models
	conn.AutoMigrate(&Mahasiswa{})
	conn.AutoMigrate(&Sisfo{})

	DB = conn

	// return
	return DB
}
