package main

import (
	"fmt"
	"math"
	"strings"
)

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64(b.panjang*b.lebar+b.panjang*b.tinggi+b.lebar*b.tinggi)
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type displayable interface {
	displayInfo() string
}

func (p phone) displayInfo() string {
	colors := strings.Join(p.colors, ", ")
	return fmt.Sprintf("Phone Info\nName: %s\nBrand: %s\nYear: %d\nColors: %s\n", p.name, p.brand, p.year, colors)
}

func printBangunDatar(bangunDatar hitungBangunDatar) {
	fmt.Printf("Luas: %d\n", bangunDatar.luas())
	fmt.Printf("Keliling: %d\n", bangunDatar.keliling())
	fmt.Println()
}

func printBangunRuang(bangunRuang hitungBangunRuang) {
	fmt.Printf("Volume: %.2f\n", bangunRuang.volume())
	fmt.Printf("Luas Permukaan: %.2f\n", bangunRuang.luasPermukaan())
	fmt.Println()
}

func display(d displayable) string {
	return d.displayInfo()
}

func luasPersegi(sisi int, displayInfo bool) interface{} {
	if sisi == 0 {
		if displayInfo {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	hasil := sisi * sisi
	if displayInfo {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, hasil)
	}
	return hasil
}

func main() {
	// Soal 1
	segitiga := segitigaSamaSisi{alas: 4, tinggi: 3}
	printBangunDatar(segitiga)

	persegiPanjang := persegiPanjang{panjang: 5, lebar: 3}
	printBangunDatar(persegiPanjang)

	tabung := tabung{jariJari: 2.5, tinggi: 7.0}
	printBangunRuang(tabung)

	balok := balok{panjang: 4, lebar: 3, tinggi: 5}
	printBangunRuang(balok)

	// Soal 2
	iphone := phone{
		name:   "iPhone X",
		brand:  "Apple",
		year:   2017,
		colors: []string{"Silver", "Space Gray"},
	}

	result := display(iphone)
	fmt.Println(result)

	// Soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	// Melakukan type assertion
	prefixStr, ok1 := prefix.(string)
	angkaPertama, ok2 := kumpulanAngkaPertama.([]int)
	angkaKedua, ok3 := kumpulanAngkaKedua.([]int)

	// Memeriksa hasil type assertion
	if !ok1 || !ok2 || !ok3 {
		fmt.Println("Type assertion gagal")
		return
	}

	// Menghitung jumlah angka
	total := 0
	for _, angka := range append(angkaPertama, angkaKedua...) {
		total += angka
	}

	// Membentuk output
	output := prefixStr
	for i, angka := range append(angkaPertama, angkaKedua...) {
		if i != 0 {
			output += " + "
		}
		output += fmt.Sprintf("%d", angka)
	}
	output += fmt.Sprintf(" = %d", total)

	fmt.Println(output)
}
