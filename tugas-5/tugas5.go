package main

import (
	"fmt"
	"strings"
)

var (
	Println = fmt.Println
	Printf  = fmt.Printf
	Print   = fmt.Print
	Sprintf = fmt.Sprintf
)

func main() {

	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	Println(luas)
	Println(keliling)
	Println(volume)

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	Println(sarah)

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	Println(buahFavoritJohn)

	// Soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	tambahDataFilm := func(title, duration, genre, year string) {
		film := map[string]string{
			"title": title,
			"jam":   duration,
			"genre": genre,
			"tahun": year,
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, v := range dataFilm {
		Println(v)
	}
}

// Function Soal 1
func luasPersegiPanjang(p, l int) int {
	return p * l
}

func kelilingPersegiPanjang(p, l int) int {
	return 2 * (p + l)
}

func volumeBalok(p, l, t int) int {
	return p * l * t
}

// Function Soal 2
func introduce(name, gender, job, age string) string {
	title := ""

	if gender == "laki-laki" {
		title = "Pak"
	} else if gender == "perempuan" {
		title = "Bu"
	} else {
		title = "Undefined"
	}

	intro := Sprintf("%s %s adalah seorang %s yang berusia %s tahun", title, name, job, age)

	return intro
}

// Function Soal 3
func buahFavorit(name string, buah ...string) string {
	mengutipBuah := make([]string, len(buah))
	for i, b := range buah {
		mengutipBuah[i] = fmt.Sprintf(`"%s"`, b)
	}

	buahString := strings.Join(mengutipBuah, ", ")
	template := fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", name, buahString)
	return template
}
