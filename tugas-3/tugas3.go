package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var (
		panjangPersegiPanjang string = "8"
		lebarPersegiPanjang   string = "5"

		panjangPersegiPanjangInt, _ = strconv.Atoi(panjangPersegiPanjang)
		lebarPersegiPanjangInt, _   = strconv.Atoi(lebarPersegiPanjang)

		alasSegitiga   string = "6"
		tinggiSegitiga string = "7"

		alasSegitigaInt, _   = strconv.Atoi(alasSegitiga)
		tinggiSegitigaInt, _ = strconv.Atoi(tinggiSegitiga)

		luasPersegiPanjang     int = panjangPersegiPanjangInt * lebarPersegiPanjangInt
		kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInt + lebarPersegiPanjangInt)
		luasSegitiga           int = int(0.5 * float64(alasSegitigaInt*tinggiSegitigaInt))
	)

	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	// Soal 2
	var (
		nilaiJohn = 80
		nilaiDoe  = 50

		messageJohn = "indeksnya John adalah"
		messageDoe  = "indeksnya Doe adalah"
	)

	if nilaiJohn >= 80 {
		fmt.Printf("%s A \n", messageJohn)
	} else if nilaiJohn >= 70 {
		fmt.Printf("%s B \n", messageJohn)
	} else if nilaiJohn >= 60 {
		fmt.Printf("%s C \n", messageJohn)
	} else if nilaiJohn >= 50 {
		fmt.Printf("%s D \n", messageJohn)
	} else {
		fmt.Printf("%s E \n", messageJohn)
	}

	if nilaiDoe >= 80 {
		fmt.Printf("%s A \n", messageDoe)
	} else if nilaiDoe >= 70 {
		fmt.Printf("%s B \n", messageDoe)
	} else if nilaiDoe >= 60 {
		fmt.Printf("%s C \n", messageDoe)
	} else if nilaiDoe >= 50 {
		fmt.Printf("%s D \n", messageDoe)
	} else {
		fmt.Printf("%s E \n", messageDoe)
	}

	// Soal 3
	var (
		tanggal = 7
		bulan   = 9
		tahun   = 2002
	)

	switch bulan {
	case 1:
		fmt.Printf("%d Januari %d \n", tanggal, tahun)
	case 2:
		fmt.Printf("%d Februari %d \n", tanggal, tahun)
	case 3:
		fmt.Printf("%d Maret %d \n", tanggal, tahun)
	case 4:
		fmt.Printf("%d April %d \n", tanggal, tahun)
	case 5:
		fmt.Printf("%d Mei %d \n", tanggal, tahun)
	case 6:
		fmt.Printf("%d Juni %d \n", tanggal, tahun)
	case 7:
		fmt.Printf("%d Juli %d \n", tanggal, tahun)
	case 8:
		fmt.Printf("%d Agustus %d \n", tanggal, tahun)
	case 9:
		fmt.Printf("%d September %d \n", tanggal, tahun)
	case 10:
		fmt.Printf("%d Oktober %d \n", tanggal, tahun)
	case 11:
		fmt.Printf("%d November %d \n", tanggal, tahun)
	case 12:
		fmt.Printf("%d Desember %d \n", tanggal, tahun)
	default:
		fmt.Println("Tidak ada nama bulan selain dari angka 1-12")
	}

	// Soal 4
	switch {
	case tahun >= 1944 && tahun <= 1964:
		fmt.Println("Baby boomer")
	case tahun >= 1965 && tahun <= 1979:
		fmt.Println("Generasi X")
	case tahun >= 1980 && tahun <= 1994:
		fmt.Println("Generasi Y")
	case tahun >= 1995 && tahun <= 2015:
		fmt.Println("Generasi Z")
	default:
		fmt.Println("undefined")
	}
}
