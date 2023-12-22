package main

import "fmt"

// Struct Soal 1
type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

// Struct Soal 2
type Segitiga struct {
	alas, tinggi int
}

type Persegi struct {
	sisi int
}

type PersegiPanjang struct {
	panjang, lebar int
}

// Struct Soal 3
type Phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *Phone) AddColor(newColor string) {
	if p.colors == nil {
		p.colors = make([]string, 0)
	}
	p.colors = append(p.colors, newColor)
}

func (b Buah) PrintInfo() {
	adaBiji := "Tidak"
	if b.AdaBijinya {
		adaBiji = "Ya"
	}

	fmt.Printf("Buah: %s, Warna: %s, Ada Biji: %s, Harga: %d Rupiah\n", b.Nama, b.Warna, adaBiji, b.Harga)
}

func (p Persegi) luasPersegi() int {
	return p.sisi * p.sisi
}

func (pp PersegiPanjang) luasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

func (sg Segitiga) luasSegitiga() float64 {
	return 0.5 * float64(sg.alas*sg.tinggi)
}

// Soal 4
type movie struct {
	Title    string
	Duration int
	Genre    string
	Year     int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	filmBaru := movie{
		Title:    title,
		Duration: duration,
		Genre:    genre,
		Year:     year,
	}

	*dataFilm = append(*dataFilm, filmBaru)
}

func tampilkanDataFilm(i int, film movie) {
	fmt.Printf("%d. title: %s\n", i, film.Title)
	fmt.Printf("   duration: %d minutes\n", film.Duration)
	fmt.Printf("   genre: %s\n", film.Genre)
	fmt.Printf("   year: %d\n", film.Year)
}

func main() {
	// Soal 1
	buahNanas := Buah{"Nanas", "Kuning", false, 9000}
	buahJeruk := Buah{"Jeruk", "Oranye", true, 8000}
	buahSemangka := Buah{"Semangka", "Hijau & Merah", true, 10000}
	pisang := Buah{"Pisang", "Kuning", false, 5000}

	buahNanas.PrintInfo()
	buahJeruk.PrintInfo()
	buahSemangka.PrintInfo()
	pisang.PrintInfo()

	// Soal 2
	persegiA := Persegi{sisi: 5}
	fmt.Printf("Luas Persegi: %d\n", persegiA.luasPersegi())

	persegiPanjangA := PersegiPanjang{panjang: 6, lebar: 10}
	fmt.Printf("Luas Persegi Panjang: %d\n", persegiPanjangA.luasPersegiPanjang())

	segitigaA := Segitiga{alas: 8, tinggi: 4}
	fmt.Printf("Luas Segitiga: %.2f\n", segitigaA.luasSegitiga())

	// Soal 3
	iphone := Phone{
		name:   "iPhone X",
		brand:  "Apple",
		year:   2017,
		colors: []string{"Silver", "Space Gray"},
	}

	fmt.Println("Warna sebelum penambahan:", iphone.colors)

	iphone.AddColor("Gold")

	fmt.Println("Warna setelah penambahan:", iphone.colors)

	// Soal 4
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		tampilkanDataFilm(i+1, film)
	}
}
