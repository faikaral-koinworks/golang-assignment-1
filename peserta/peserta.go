package peserta

import "fmt"

type Person struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var Murid = []Person{
	{nama: "A", alamat: "Jakarta", pekerjaan: "Back-end engineer", alasan: "golang is cool"},
	{nama: "B", alamat: "Tangerang", pekerjaan: "Back-end engineer", alasan: "golang is good"},
	{nama: "C", alamat: "Bogor", pekerjaan: "Back-end engineer", alasan: "golang is awesome"},
	{nama: "D", alamat: "Bekasi", pekerjaan: "Back-end engineer", alasan: "golang is interesting"},
	{nama: "E", alamat: "Depok", pekerjaan: "Back-end engineer", alasan: "golang is easy"},
	{nama: "F", alamat: "Serang", pekerjaan: "Back-end engineer", alasan: "golang is fast"},
}

func (p Person) Cetak() {
	fmt.Printf("Nama: %v\n", p.nama)
	fmt.Printf("Alamat: %v\n", p.alamat)
	fmt.Printf("Pekerjaan: %v\n", p.pekerjaan)
	fmt.Printf("Alasan memilih kelas golang: %v\n", p.alasan)
}
