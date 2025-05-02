package database

type Penulis struct {
    ID int
    Nama string
}

type Ilustrator struct {
    ID int
    Nama string
}

type Penerbit struct {
    ID int
    Nama string
}

type Genre struct {
    ID int
    Nama string
}

type Komik struct {
    ID int
    Judul string
    Penulis Penulis
    Ilustrator Ilustrator
    Genre Genre
    TahunTerbit int
    Penerbit Penerbit
    Status string
    Rating float64
}