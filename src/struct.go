/* 
-- Pendefinisian struct ada pada file ini
*/

package src

/* 
-- Disini memakai 2 struct, yaitu Person dan Student
*/


// Struct Person mempunyai field Nama, Alamat, dan Pekerjaan
type Person struct {
	Nama string
	Alamat string
	Pekerjaan string
}

/* Struct Student mempunyai field yang ada di Person ditambah
Absen dan Alasan */
type Student struct {
	person Person	
	Absen int
	Alasan string
}