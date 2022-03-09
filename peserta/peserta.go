package peserta

import "fmt"

type Person struct {
	absen     int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var Murid = []Person{
	{absen: 1, nama: "A", alamat: "Jakarta", pekerjaan: "Back-end engineer", alasan: "golang is cool"},
	{absen: 2, nama: "B", alamat: "Tangerang", pekerjaan: "Back-end engineer", alasan: "golang is good"},
	{absen: 3, nama: "C", alamat: "Bogor", pekerjaan: "Back-end engineer", alasan: "golang is awesome"},
	{absen: 4, nama: "D", alamat: "Bekasi", pekerjaan: "Back-end engineer", alasan: "golang is interesting"},
	{absen: 5, nama: "E", alamat: "Depok", pekerjaan: "Back-end engineer", alasan: "golang is easy"},
	{absen: 6, nama: "F", alamat: "Serang", pekerjaan: "Back-end engineer", alasan: "golang is fast"},
}

func (p Person) Cetak() {
	fmt.Printf("Nama: %v\n", p.nama)
	fmt.Printf("Alamat: %v\n", p.alamat)
	fmt.Printf("Pekerjaan: %v\n", p.pekerjaan)
	fmt.Printf("Alasan memilih kelas golang: %v\n", p.alasan)
}
