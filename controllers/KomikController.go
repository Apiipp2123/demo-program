package controller

import (
	"fmt"
	"node/database"
	"strings"
)

var KomikList []database.Komik
var komikIDCounter = 1
var entityIDCounter = 1

func TambahKomik(
	judul, namaPenulis, namaIlustrator, namaGenre string,
	tahunTerbit int,
	namaPenerbit string,
	status string,
	rating float64,
) {
	komik := database.Komik{
		ID: komikIDCounter,
		Judul: strings.TrimSpace(judul),
		Penulis: database.Penulis{ID: entityIDCounter, Nama: strings.TrimSpace(namaPenulis)},
		Ilustrator: database.Ilustrator{ID: entityIDCounter + 1, Nama: strings.TrimSpace(namaIlustrator)},
		Genre: database.Genre{ID: entityIDCounter + 2, Nama: strings.TrimSpace(namaGenre)},
		TahunTerbit: tahunTerbit,
		Penerbit: database.Penerbit{ID: entityIDCounter + 3, Nama: strings.TrimSpace(namaPenerbit)},
		Status: strings.TrimSpace(status),
		Rating: rating,
	}
	KomikList = append(KomikList, komik)
	komikIDCounter++
	entityIDCounter += 4
}

func TampilkanSemuaKomik() {
	if len(KomikList) == 0 {
		fmt.Println("Belum ada data komik.")
		return
	}
	for _, k := range KomikList {
		fmt.Println("ID: ", k.ID)
		fmt.Println("Judul: ", k.Judul)
		fmt.Println("Penulis: ", k.Penulis.Nama)
		fmt.Println("Genre: ", k.Genre.Nama)
		fmt.Println("Rating: ", k.Rating)
	}
}

func UpdateKomik(id int, judulBaru string) {
	for i := range KomikList {
		if KomikList[i].ID == id {
			KomikList[i].Judul = strings.TrimSpace(judulBaru)
			fmt.Println("Komik berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Komik tidak ditemukan.")
}

func DeleteKomik(id int) {
	for i := range KomikList {
		if KomikList[i].ID == id {
			KomikList = append(KomikList[:i], KomikList[i+1:]...)
			fmt.Println("Komik berhasil dihapus.")
			return
		}
	}
	fmt.Println("Komik tidak ditemukan.")
}

func CariKomikByJudul(keyword string) {
	keyword = strings.ToLower(keyword)
	found := false
	for _, k := range KomikList {
		if strings.Contains(strings.ToLower(k.Judul), keyword) {
			fmt.Println("ID: ", k.ID)
		fmt.Println("Judul: ", k.Judul)
		fmt.Println("Penulis: ", k.Penulis.Nama)
		fmt.Println("Genre: ", k.Genre.Nama)
		fmt.Println("Rating: ", k.Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

func CariKomikByGenre(keyword string) {
	keyword = strings.ToLower(keyword)
	found := false
	for _, k := range KomikList {
		if strings.Contains(strings.ToLower(k.Genre.Nama), keyword) {
			fmt.Println("ID: ", k.ID)
		fmt.Println("Judul: ", k.Judul)
		fmt.Println("Penulis: ", k.Penulis.Nama)
		fmt.Println("Genre: ", k.Genre.Nama)
		fmt.Println("Rating: ", k.Rating)
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}
