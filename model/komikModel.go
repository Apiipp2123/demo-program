package model

import (
	"node/database"
	"strings"
)

func NewKomik(id, tahunTerbit, jumlahHalaman int, judul, penulis, ilustrator, genre, penerbit, status string, rating float64) database.Komik {
	return database.Komik{
		ID:    id,
		Judul: strings.TrimSpace(judul),
		Penulis: database.Penulis{
			ID:   id * 10,
			Nama: strings.TrimSpace(penulis),
		},
		Ilustrator: database.Ilustrator{
			ID:   id * 10 + 1,
			Nama: strings.TrimSpace(ilustrator),
		},
		Genre: database.Genre{
			ID:   id * 10 + 2,
			Nama: strings.TrimSpace(genre),
		},
		TahunTerbit: tahunTerbit,
		Penerbit: database.Penerbit{
			ID:   id * 10 + 3,
			Nama: strings.TrimSpace(penerbit),
		},
	}
}
