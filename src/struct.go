package src

type Person struct {
	Nama string
	Alamat string
	Pekerjaan string
}

type Student struct {
	person Person	
	Absen int
	Alasan string
}