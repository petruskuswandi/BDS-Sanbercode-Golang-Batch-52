package shapes

import (
	"fmt"
	"math"
	"strings"
)

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64(b.Panjang*b.Lebar+b.Panjang*b.Tinggi+b.Lebar*b.Tinggi)
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type Displayable interface {
	DisplayInfo() string
}

func (p Phone) DisplayInfo() string {
	colors := strings.Join(p.Colors, ", ")
	return fmt.Sprintf("Phone Info\nName: %s\nBrand: %s\nYear: %d\nColors: %s\n", p.Name, p.Brand, p.Year, colors)
}

func LuasPersegi(sisi int, displayInfo bool) interface{} {
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
