package models

// buat struct mahasiswa sebagai model atau representasi data mahasiswa dari database nantinya
type Mahasiswa struct {
	Nim     string `json:"nim"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}
