package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 2
	var sks = 126
	var isKampusMerdeka = !true

	var kkmNilaiAkhir = 80
	var maxAbsensi = 3
	var maxSks = 126

	var lulus = nilaiAkhir >= kkmNilaiAkhir && absensi < maxAbsensi && sks == maxSks || isKampusMerdeka

	fmt.Println(lulus)
}
