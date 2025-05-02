package view

import (
	"bufio"
	"fmt"
	controller "node/controllers"
	"os"
	"strconv"
	"strings"
)

func Inputkomik() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul: ")
	judul, _ := reader.ReadString('\n')

	fmt.Print("Penulis: ")
	penulis, _ := reader.ReadString('\n')

	fmt.Print("Ilustrator: ")
	ilustrator, _ := reader.ReadString('\n')

	fmt.Print("Genre: ")
	genre, _ := reader.ReadString('\n')

	fmt.Print("Tahun Terbit: ")
	tahunStr, _ := reader.ReadString('\n')
	tahun, _ := strconv.Atoi(strings.TrimSpace(tahunStr))

	fmt.Print("Penerbit: ")
	penerbit, _ := reader.ReadString('\n')

	fmt.Print("Status: ")
	status, _ := reader.ReadString('\n')

	fmt.Print("Rating: ")
	ratingStr, _ := reader.ReadString('\n')
	rating, _ := strconv.ParseFloat(strings.TrimSpace(ratingStr), 64)

	controller.TambahKomik(judul, penulis, ilustrator, genre, tahun, penerbit, status, rating)
}
