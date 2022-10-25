package peserta

import "fmt"

type Person struct {
	absen     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var Murid = []Person{
	{absen: 1, nama: "Septian", alamat: "Subang", pekerjaan: "Admin", alasan: "Its Free"},
	{absen: 2, nama: "Asep", alamat: "Bandung", pekerjaan: "Sales", alasan: "Golang Mantep"},
	{absen: 3, nama: "Lingling", alamat: "Bogor", pekerjaan: "IT Support", alasan: "Menarik untuk dipelajari"},
	{absen: 4, nama: "Fatur", alamat: "Bekasi", pekerjaan: "Pelajar", alasan: "Ingin menjadi programmer"},
	{absen: 5, nama: "Zilong", alamat: "Tanggerang", pekerjaan: "Pelajar", alasan: "Hacktiv8 Mantap"},
}

func (p Person) Cetak() {
	fmt.Printf("Nama: %v\n", p.nama)
	fmt.Printf("Alamat: %v\n", p.alamat)
	fmt.Printf("Pekerjaan: %v\n", p.pekerjaan)
	fmt.Printf("Alasan memilih kelas golang: %v\n", p.alasan)
}
