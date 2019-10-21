package PMB

import (
	"fmt"
	"sort"
	"strings"
)

type peserta struct {
	id int
	nama string
	nilai[4] float64
	total float64
}

var pesertaList[] peserta
var index = 1

func TambahPeserta(namePeserta string)  {
	p := peserta{nama: namePeserta, id:index}
	pesertaList = append(pesertaList,p)
	index++
}

func TambahNilaiPeserta(searchID int, nilai1 float64, nilai2 float64, nilai3 float64, nilai4 float64){
	if len(pesertaList) == 0 {
		fmt.Println("Masih Belum ada Data")
	} else {
		search := false
		for i:= 0; i<len(pesertaList); i++ {
			if pesertaList[i].id == searchID {
				pesertaList[i].nilai = [4]float64{nilai1,nilai2,nilai3,nilai4}
				pesertaList[i].total = nilai1 + nilai2 + nilai3 + nilai4
				search = true
				break
			}
		}
		if search==false {
			fmt.Println("ID Tidak ditemukan")
		}
	}
}

func SearchName(name string){
	search := false
	fmt.Println("Cari Nama")
	for i:=0; i<len(pesertaList); i++ {
		if strings.Contains(pesertaList[i].nama,name) {
			search = true
			fmt.Println("Id :",pesertaList[i].id)
			fmt.Println("Nama :",pesertaList[i].nama)
			fmt.Println("Nilai :",pesertaList[i].nilai)
			fmt.Println("Total Nilai :",pesertaList[i].total)
			fmt.Println()
		}
	}
	if search==false {
		fmt.Println("Nama Tidak Ditemukan")
	}
}

func View(){
	if len(pesertaList) == 0 {
		fmt.Println("Masih Belum ada Data")
	} else {
		sort.Slice(pesertaList, func(i, j int) bool {
			return pesertaList[i].id < pesertaList[j].id
		})
		fmt.Println("View Data")
		for i:= 0; i<len(pesertaList); i++ {
			fmt.Println("Id :",pesertaList[i].id)
			fmt.Println("Nama :",pesertaList[i].nama)
			fmt.Println("Nilai :",pesertaList[i].nilai)
			fmt.Println("Total Nilai :",pesertaList[i].total)
			fmt.Println()
		}
	}
}

func TopThree()  {
	if len(pesertaList) == 0 {
		fmt.Println("Masih Belum ada Data")
	} else {
		sort.Slice(pesertaList, func(i, j int) bool {
			return pesertaList[i].total > pesertaList[j].total
		})
		fmt.Println("Top Three")
		if len(pesertaList) <= 2 {
			for i:=0; i<len(pesertaList); i++ {
				fmt.Println("Id :",pesertaList[i].id)
				fmt.Println("Nama :",pesertaList[i].nama)
				fmt.Println("Nilai :",pesertaList[i].nilai)
				fmt.Println("Total Nilai :",pesertaList[i].total)
				fmt.Println()
			}
		} else {
			for i:=0; i<3; i++ {
				fmt.Println("Id :",pesertaList[i].id)
				fmt.Println("Nama :",pesertaList[i].nama)
				fmt.Println("Nilai :",pesertaList[i].nilai)
				fmt.Println("Total Nilai :",pesertaList[i].total)
				fmt.Println()
			}
		}
	}
}