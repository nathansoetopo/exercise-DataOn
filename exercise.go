package main

import "fmt"

type Roda interface {
	JenisBan() string
	spin()
}

type BanKaret struct{}

func (b BanKaret) JenisBan() string {
	return "Ban Karet"
}

func (b BanKaret) spin() {}

type BanKayu struct{}

func (b BanKayu) JenisBan() string {
	return "Ban Kayu"
}

func (b BanKayu) spin() {}

type BanBesi struct{}

func (b BanBesi) JenisBan() string {
	return "Ban Besi"
}

func (b BanBesi) spin() {}

type Pintu interface {
	Ketuk() string
	Buka() string
}

type PintuKanan struct{}

func (p PintuKanan) Ketuk() string {
	return "Knock"
}

func (p PintuKanan) Buka() string {
	return "Beep"
}

// PintuKiri adalah pintu kiri mobil
type PintuKiri struct{}

func (p PintuKiri) Ketuk() string {
	return "Beep"
}

func (p PintuKiri) Buka() string {
	return "Knock"
}

// Mobil adalah struktur untuk merepresentasikan sebuah mobil
type Mobil struct {
	Roda       [4]Roda
	PintuKanan Pintu
	PintuKiri  Pintu
}

// GantiRoda mengganti jenis roda pada mobil
func (mobil *Mobil) GantiRoda(rodaDepanKanan, rodaDepanKiri, rodaBelakangKanan, rodaBelakangKiri Roda) {
	mobil.Roda[0] = rodaDepanKanan
	mobil.Roda[1] = rodaDepanKiri
	mobil.Roda[2] = rodaBelakangKanan
	mobil.Roda[3] = rodaBelakangKiri
}

func main() {
	// Membuat objek mobil
	mobilSaya := Mobil{}
	pintuKanan := PintuKanan{}
	pintuKiri := PintuKiri{}

	// Mengganti roda mobil dengan berbagai jenis ban
	mobilSaya.GantiRoda(BanKaret{}, BanKayu{}, BanBesi{}, BanKaret{})

	// Menggunakan pintu kanan
	fmt.Println("Ketuk pintu kanan:", pintuKanan.Ketuk()) // Output: "Knock"
	fmt.Println("Buka pintu kanan:", pintuKanan.Buka())   // Output: "Beep"

	// Menggunakan pintu kiri
	fmt.Println("Ketuk pintu kiri:", pintuKiri.Ketuk()) // Output: "Beep"
	fmt.Println("Buka pintu kiri:", pintuKiri.Buka())   // Output: "Knock"
}
