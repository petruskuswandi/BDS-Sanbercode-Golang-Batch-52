package main

import "fmt"

func main() {
	// Soal 1
	for i := 1; i <= 20; i++ {
		if i%2 != 0 && i%3 != 0 {
			fmt.Printf("%d - Santai \n", i)
		}
		if i%2 == 0 {
			fmt.Printf("%d - Berkualitas \n", i)
		}
		if i%3 == 0 && i%2 != 0 {
			fmt.Printf("%d - I Love Coding \n", i)
		}
	}

	// Soal 2
	for j := 1; j <= 7; j++ {
		fmt.Println()
		for k := 1; k <= j; k++ {
			fmt.Print("#")
		}
	}

	fmt.Println()
	fmt.Println()
	// Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])
	fmt.Println("")

	// Soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for i, v := range sayuran {
		fmt.Printf("%d. %s \n", i+1, v)
	}
	fmt.Println()

	// Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	var volumeBalok = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	fmt.Printf("Panjang = %d \n", satuan["panjang"])
	fmt.Printf("Lebar = %d \n", satuan["lebar"])
	fmt.Printf("Tinggi = %d \n", satuan["tinggi"])

	// for k, v := range satuan {
	// 	fmt.Printf("%s = %d \n", k, v)
	// }

	fmt.Printf("Volume Balok = %d \n", volumeBalok)
}
