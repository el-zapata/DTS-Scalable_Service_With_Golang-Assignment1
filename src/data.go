/* 
-- Data setiap siswa disimpan dalam file ini 
-- Data berupa struct dan disimpan dalam slice (Students)
*/


package src

var Students = []Student{
	{	person: Person{	Nama: "Zapata", 
						Alamat: "Jalan Batas I", 
						Pekerjaan: "Mahasiswa",
					},
		Absen: 1,
		Alasan: "Coba-coba",
	},

	{	person: Person{	Nama: "Dika", 
						Alamat: "Jalan Batas I", 
						Pekerjaan: "Siswa",
					},
		Absen: 2,
		Alasan: "Suka dengan Golang",
	},

	{	person: Person{	Nama: "Dahana", 
						Alamat: "Jalan Batas I", 
						Pekerjaan: "Siswa",
					},
		Absen: 3,
		Alasan: "Saya dipaksa",
	},

	{	person: Person{	Nama: "Intan", 
						Alamat: "Jalan Batas I", 
						Pekerjaan: "Siswa",
					},
		Absen: 4,
		Alasan: "Ntah, saya hanya disuruh untuk mengikuti ini",
	},

	{	person: Person{	Nama: "Alvent", 
						Alamat: "Jalan Kejora", 
						Pekerjaan: "Pegawai Negeri Sipil",
					},
		Absen: 5,
		Alasan: "Memperkaya ilmu IT saya",
	},

	{	person: Person{	Nama: "Rara", 
						Alamat: "Jalan Teratai", 
						Pekerjaan: "Guru",
					},
		Absen: 6,
		Alasan: "Mengisi waktu luang",
	},

	{	person: Person{	Nama: "Miko", 
						Alamat: "Jalan Gerobak", 
						Pekerjaan: "Wiraswasta",
					},
		Absen: 7,
		Alasan: "Ingin mencari jodoh dengan orang-orang IT",
	},

	{	person: Person{	Nama: "Mika", 
						Alamat: "Jalan Gerobak", 
						Pekerjaan: "Pilot",
					},
		Absen: 8,
		Alasan: "Capek terbang terus, mau belajar back-end",
	},

	{	person: Person{	Nama: "Mimo", 
						Alamat: "Jalan Ninja III", 
						Pekerjaan: "Ninja",
					},
		Absen: 9,
		Alasan: "Mau cari profesi lain gara-gara dimarahin gagal misi terus",
	},

	{	person: Person{	Nama: "Mima", 
						Alamat: "Jalan Setir", 
						Pekerjaan: "Pembalap",
					},
		Absen: 10,
		Alasan: "Mau banting setir jadi programmer",
	},
}
