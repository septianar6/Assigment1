package main

import (
	"biodata/peserta"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please enter Args")
		return
	}
	absen, err := strconv.Atoi(os.Args[1])
	_ = err

	if absen > len(peserta.Murid) || absen < 1 {
		fmt.Println("Invalid Absen")
		return
	}

	peserta.Murid[absen-1].Cetak()
}
