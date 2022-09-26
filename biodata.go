package main

import (
	"Assignment_1/src" // Import file struct, data, dan cek_absen
	"fmt" // Untuk print
	"os" // Untuk os.Args (Argument pada command line)
	"strconv" // Untuk convert string ke integer (Atoi)
)

/*
-- Nilai os.Args diambil yang elemen ke 1 (dimana nilai argument pada command line berada)
-- key_input merupakan pointer string (Untuk menghemat memori?)
*/
var key_input *string = &os.Args[1]

func main() {

	// Variabel absen untuk nilai hasil konversi, dan err jika terjadi error pada proses konversi
	absen, err := strconv.Atoi(*key_input)

	if err != nil { 
		// Jika err tidak nil (Terjadi error)
		fmt.Printf("Tolong berikan input yang valid")
	} else { 
		// Jika tidak terjadi error pada konversi
		fmt.Printf("Berikut data siswa absen %d\n", absen)
		// Print hasil return dari fungsi Cek_absen (dengan passing alamat memori absen)
		// Digunakan alamat memori absen karena fungsi ini menerima input pointer
		fmt.Print(src.Cek_absen(&absen))
	}
}