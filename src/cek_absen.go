package src

import (
	"fmt"
)

// Fungsi Cek_absen mengambil input parameter pointer dan nilai return bertipe string
func Cek_absen(input *int) string {
	data := "Data siswa tidak ditemukan!!!"
	outerLoop:
	for _, data_student := range Students {
		for range []interface{} {	data_student.person.Nama, 
									data_student.person.Alamat,
									data_student.person.Pekerjaan,
									data_student.Absen, 
									data_student.Alasan,
									} {
		if (*input == data_student.Absen) {
			data = fmt.Sprintf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n",
								data_student.person.Nama,
								data_student.person.Alamat,
								data_student.person.Pekerjaan,
								data_student.Alasan)

			// Akan break dari seluruh for loop sehingga tidak perlu cek data selanjutnya
			break outerLoop 
		}
		}
	}
	return data
}

/* Alasan saya memakai empty interface adalah sebagai berikut:

Tadinya saya mau iterate seluruh field di dalam satu data dengan:

	for _, data_student := range Students {
		for field, value := range data_student {
			
		}
	}
	
Namun error: Cannot range over data_student (variable of type Student)

Lalu saya melihat di stackoverflow bahwa untuk iterate value dari
field bisa digunakan empty interface seperti di atas
*/

