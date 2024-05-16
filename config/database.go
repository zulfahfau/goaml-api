package config

import (
	"log" // Ini buat logging

	"gorm.io/driver/postgres" // Ini driver untuk PostgreSQL
	"gorm.io/gorm"            // Ini GORM, orm untuk Golang
)

// DB buat nyimpen koneksi database
var DB *gorm.DB

// ConnectDB buat ngelakuin koneksi ke database
func ConnectDB() {
	dsn := "host=localhost user=goaml_user password=password dbname=goaml port=5432 sslmode=disable" // Nah, ini string koneksi ke database
	var err error
	// Ini buat nyambungin ke database dengan GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// Kalo gagal nyambung, ya udah, langsung exit
	if err != nil {
		log.Fatal("Gagal konek database:", err) // Udah, kalo gagal langsung keluar aja
	}
	log.Println("Database terkoneksi") // Yoi, kalo berhasil nyambung, log ini biar tahu
}
