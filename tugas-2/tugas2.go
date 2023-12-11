package main

import (
	"fmt"
	"strconv"
	"strings"
)

var Println = fmt.Println

func main() {
	// Soal 2
	halo := "Halo Dunia"

	halo = strings.Replace(halo, "Dunia", "Golang", -1)
	Println(halo)

	// Soal 3
	var (
		kataPertama = "saya"
		kataKedua   = "senang"
		kataKetiga  = "belajar"
		kataKeempat = "golang"
	)
	var capslockKataKeempat = strings.ToUpper(kataKeempat)

	fmt.Printf("%s %s %s %s \n", kataPertama, kataKedua, kataKetiga, capslockKataKeempat)

	// Soal 4
	var (
		angkaPertama       = "8"
		angkaKedua         = "5"
		angkaKetiga        = "6"
		angkaKeempat       = "7"
		angkaPertamaInt, _ = strconv.Atoi(angkaPertama)
		angkaKeduaInt, _   = strconv.Atoi(angkaKedua)
		angkaKetigaInt, _  = strconv.Atoi(angkaKetiga)
		angkaKeempatInt, _ = strconv.Atoi(angkaKeempat)
	)

	fmt.Println(angkaPertamaInt + angkaKeduaInt + angkaKetigaInt + angkaKeempatInt)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	changeKalimat := strings.Replace(kalimat, "halo halo", "hi hi", -1)
	fmt.Println(changeKalimat, "-", angka)
}
