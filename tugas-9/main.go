package main

import (
	"fmt"
	"tugas-9/shapes"
)

func main() {
	// Soal 1
	segitiga := shapes.SegitigaSamaSisi{Alas: 4, Tinggi: 3}
	printBangunDatar(segitiga)

	persegiPanjang := shapes.PersegiPanjang{Panjang: 5, Lebar: 3}
	printBangunDatar(persegiPanjang)

	tabung := shapes.Tabung{JariJari: 2.5, Tinggi: 7.0}
	printBangunRuang(tabung)

	balok := shapes.Balok{Panjang: 4, Lebar: 3, Tinggi: 5}
	printBangunRuang(balok)

	// Soal 2
	iphone := shapes.Phone{
		Name:   "iPhone X",
		Brand:  "Apple",
		Year:   2017,
		Colors: []string{"Silver", "Space Gray"},
	}

	result := display(iphone)
	fmt.Println(result)

	// Soal 3
	fmt.Println(shapes.LuasPersegi(4, true))
	fmt.Println(shapes.LuasPersegi(8, false))
	fmt.Println(shapes.LuasPersegi(0, true))
	fmt.Println(shapes.LuasPersegi(0, false))

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

func printBangunDatar(bangunDatar shapes.HitungBangunDatar) {
	fmt.Printf("Luas: %d\n", bangunDatar.Luas())
	fmt.Printf("Keliling: %d\n", bangunDatar.Keliling())
	fmt.Println()
}

func printBangunRuang(bangunRuang shapes.HitungBangunRuang) {
	fmt.Printf("Volume: %.2f\n", bangunRuang.Volume())
	fmt.Printf("Luas Permukaan: %.2f\n", bangunRuang.LuasPermukaan())
	fmt.Println()
}

func display(d shapes.Displayable) string {
	return d.DisplayInfo()
}
