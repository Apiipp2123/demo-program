package view

import (
	"bufio"
	"fmt"
	controller "node/controllers"
	"os"
	"strconv"
	"strings"
)

func OpsiTampilan() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Aplikasi Data Komik CLI ===")
		fmt.Println("1. Tambah Komik")
		fmt.Println("2. Tampilkan Semua Komik")
		fmt.Println("3. Update Komik (Judul)")
		fmt.Println("4. Hapus Komik")
		fmt.Println("5. Cari Komik (Judul)")
		fmt.Println("6. Cari Komik (Genre)")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			Inputkomik()
		case "2":
			controller.TampilkanSemuaKomik()
		case "3":
			UpdateKomikView()
		case "4":
			DeleteKomikView()
		case "5":
			CariJudulView()
		case "6":
			CariGenreView()
		case "7":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Println()
	}
}

func UpdateKomikView() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID komik yang diupdate: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Judul baru: ")
	judul, _ := reader.ReadString('\n')
	controller.UpdateKomik(id, judul)
}

func DeleteKomikView() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID komik yang dihapus: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	controller.DeleteKomik(id)
}

func CariJudulView() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Cari judul: ")
	judul, _ := reader.ReadString('\n')
	controller.CariKomikByJudul(strings.TrimSpace(judul))
}

func CariGenreView() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Cari genre: ")
	genre, _ := reader.ReadString('\n')
	controller.CariKomikByGenre(strings.TrimSpace(genre))
}
