package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// Function Soal 1
func contoh(sentence string, year int) {
	defer func() {
		fmt.Printf("Kalimat: %s, Tahun: %d\n", sentence, year)
	}()

	fmt.Println("Kalimat yang didahulukan")
}

// Function Soal 2
func kelilingSegitigaSamaSisi(sisi int, printString bool) (string, error) {
	if sisi == 0 {
		err := errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
		if printString {
			return "", err
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic recovered:", r)
				}
			}()

			panic(err)
		}
	}
	keliling := sisi * 3
	if printString {
		return fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	} else {
		return fmt.Sprintf("%d", keliling), nil
	}
}

// Function Soal 3
func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Printf("Total angka: %d\n", *total)
}

// Function Soal 4
func tambahPhone(phones *[]string, brand string) {
	*phones = append(*phones, brand)
}

// Function Soal 5
func luasLingkaran(r float64) int {
	luas := math.Pi * math.Pow(r, 2)
	return int(math.Round(luas))
}

func kelilingLingkaran(r float64) int {
	keliling := 2 * math.Pi * r
	return int(math.Round(keliling))
}

// Function Soal 6
func luasPersegiPanjang(p, l float64) float64 {
	return p * l
}

func kelilingPersegiPanjang(p, l float64) float64 {
	return 2 * (p + l)
}

func main() {
	// Soal 1
	defer fmt.Println("Program selesai.")

	kalimat := "Golang Backend Development"
	tahun := 2021
	contoh(kalimat, tahun)

	// Soal 2
	// Contoh pemanggilan fungsi dengan parameter true
	result, err := kelilingSegitigaSamaSisi(4, true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Contoh pemanggilan fungsi dengan parameter false
	result, err = kelilingSegitigaSamaSisi(8, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Contoh pemanggilan fungsi dengan parameter 0 dan true
	result, err = kelilingSegitigaSamaSisi(0, true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// Contoh pemanggilan fungsi dengan parameter 0 dan false
	_, err = kelilingSegitigaSamaSisi(0, false)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Soal 3
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 4
	var phones = []string{}

	tambahPhone(&phones, "Xiaomi")
	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "IPhone")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Vivo")

	sort.Strings(phones)

	for i, v := range phones {
		nomorUrut := i + 1

		time.Sleep(time.Second)
		fmt.Printf("%d. %s\n", nomorUrut, v)
	}

	// Soal 5
	r1 := 7.0
	r2 := 10.0
	r3 := 15.0

	luas1 := luasLingkaran(r1)
	keliling1 := kelilingLingkaran(r1)

	luas2 := luasLingkaran(r2)
	keliling2 := kelilingLingkaran(r2)

	luas3 := luasLingkaran(r3)
	keliling3 := kelilingLingkaran(r3)

	fmt.Printf("Jari-Jari: %.1f, Luas: %d, keliling: %d\n", r1, luas1, keliling1)
	fmt.Printf("Jari-Jari: %.1f, Luas: %d, keliling: %d\n", r2, luas2, keliling2)
	fmt.Printf("Jari-Jari: %.1f, Luas: %d, keliling: %d\n", r3, luas3, keliling3)

	// Soal 6
	var panjang, lebar float64

	flag.Float64Var(&panjang, "panjang", 0.0, "Panjang persegi Panjang")
	flag.Float64Var(&lebar, "lebar", 0.0, "Lebar persegi panjang")
	flag.Parse()

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)

	fmt.Printf("Luas persegi panjang dengan panjang %.2f dan lebar %.2f adalah %.2f\n", panjang, lebar, luas)
	fmt.Printf("Keliling persegi panjang dengan panjang %.2f dan lebar %.2f adalah %.2f\n", panjang, lebar, keliling)
}
