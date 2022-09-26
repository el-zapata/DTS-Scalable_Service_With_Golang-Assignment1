package src

import (
	"fmt"
)

/*
-- Fungsi Cek_absen mengambil input parameter pointer dan nilai return bertipe string
-- Fungsi ini akan mengembalikan string variabel data, jika no absen ditemukan, data akan
bernilai data dari siswa absen tersebut. Jika tidak, akan bernilai "Data siswa tidak
ditemukan!!!""
*/ 

func Cek_absen(input *int) string {
	data := "Data siswa tidak ditemukan!!!"
	for _, data_student := range Students {
		if (*input == data_student.Absen) {
			data = fmt.Sprintf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n",
								data_student.person.Nama,
								data_student.person.Alamat,
								data_student.person.Pekerjaan,
								data_student.Alasan)

			// Akan break dari seluruh for loop sehingga tidak perlu cek data selanjutnya
			break
		}
	}
	return data
}
