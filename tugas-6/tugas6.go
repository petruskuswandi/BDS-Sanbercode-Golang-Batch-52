package main

import (
	"fmt"
	"math"
)

var (
	// Soal 1
	luasLingkaran     float64
	kelilingLingkaran float64
	Printf            = fmt.Printf
	Println           = fmt.Println
	Print             = fmt.Print
	Sprintf           = fmt.Sprintf
	// Soal 2
	sentence string

	// Soal 3
	buah = []string{}

	// Soal 4
	dataFilm = []map[string]string{}
)

const Pi = 3.14

func main() {
	// Soal 1
	r := 5.0

	hitungLingkaran(&r)

	Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	// Soal 2
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	result := addBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
	Print(result)

	// Soal 4
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		tampilkanDataFilm(i+1, film)
	}
}

// Function Soal 1
func hitungLingkaran(r *float64) {
	luasLingkaran = math.Pi * *r * *r
	kelilingLingkaran = 2 * Pi * *r
}

// Function Soal 2
func introduce(sentence *string, name, gender, job, age string) {
	panggilan := "Pak"
	if gender == "perempuan" {
		panggilan = "Bu"
	}

	*sentence = Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggilan, name, job, age)
}

// Function Soal 3
func addBuah(buah *[]string, data ...string) string {
	*buah = append(*buah, data...)

	result := ""
	for i, v := range *buah {
		result += Sprintf("%d. %s \n", i+1, v)
	}

	return result
}

// Function Soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func tampilkanDataFilm(i int, film map[string]string) {
	fmt.Printf("%d. title: %s\n", i, film["title"])
	fmt.Printf("   duration: %s\n", film["duration"])
	fmt.Printf("   genre: %s\n", film["genre"])
	fmt.Printf("   year: %s\n", film["year"])
}
