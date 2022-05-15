package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/restfull-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// cek error
	if err != nil {
		log.Panicln("Gagal terhubung ke database")
		panic(err)
	}

	// jika berhasil maka lakukan migration models
	db.AutoMigrate(&Mahasiswa{})

	// return
	return db
}
