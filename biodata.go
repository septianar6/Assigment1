package main

import (
	"fmt"
	"os"
)

type Anggota struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	absen := os.Args[1]
	var s1 = Anggota{"Septian", "Subang", "Admin", "Its Free"}
	var s2 = Anggota{"Asep", "Bandung", "Sales", "Golang Mantap"}
	var s3 = Anggota{"Lingling", "Subang", "IT Support", "Menarik untuk dipelajari"}
	var s4 = Anggota{"Fatur", "Jakarta", "Pelajar", "Ingin menjadi Programmer"}
	var s5 = Anggota{"Zilong", "Tanggerang", "Pelajar", "Hacktiv8 Mantap"}

	if absen == "1" {
		fmt.Println("Nama :", s1.nama)
		fmt.Println("Alamat :", s1.alamat)
		fmt.Println("Pekerjaan :", s1.pekerjaan)
		fmt.Println("Alasan :", s1.alasan)
	} else if absen == "2" {
		fmt.Println("Nama :", s2.nama)
		fmt.Println("Alamat :", s2.alamat)
		fmt.Println("Pekerjaan :", s2.pekerjaan)
		fmt.Println("Alasan :", s2.alasan)
	} else if absen == "3" {
		fmt.Println("Nama :", s3.nama)
		fmt.Println("Alamat :", s3.alamat)
		fmt.Println("Pekerjaan :", s3.pekerjaan)
		fmt.Println("Alasan :", s3.alasan)
	} else if absen == "4" {
		fmt.Println("Nama :", s4.nama)
		fmt.Println("Alamat :", s4.alamat)
		fmt.Println("Pekerjaan :", s4.pekerjaan)
		fmt.Println("Alasan :", s4.alasan)
	} else if absen == "5" {
		fmt.Println("Nama :", s5.nama)
		fmt.Println("Alamat :", s5.alamat)
		fmt.Println("Pekerjaan :", s5.pekerjaan)
		fmt.Println("Alasan :", s5.alasan)
	} else {
		fmt.Println("Hanya Memiliki 5 Data!")
	}
}
