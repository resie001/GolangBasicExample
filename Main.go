package main

import (
	"awesomeProject/PMB"
	"fmt"
)

func main() {
	choice := 0
	for {
		fmt.Println("Menu")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Tambah Nilai Peserta")
		fmt.Println("3. Search By Nama")
		fmt.Println("4. Lihat Semua Peserta")
		fmt.Println("5. Lihat Top Three")
		fmt.Println("6. Exit")
		fmt.Print("Masukin Pilihan: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1 :
			addParticipant()
		case 2:
			addScoreParticipants()
		case 3:
			searchParticipantByName()
		case 4:
			PMB.View()
		case 5:
			PMB.TopThree()
		case 6:
			fmt.Println("Exiting App....")
		default:
			fmt.Println("Out of Ranges!!!")
		}
		if choice == 6 {
			break
		}
	}
}

func addParticipant(){
	fmt.Println("Tambah Peserta")
	fmt.Print("Nama Peserta: ")
	namaPeserta := ""
	fmt.Scanln(&namaPeserta)
	PMB.TambahPeserta(namaPeserta)
}

func addScoreParticipants(){
	fmt.Println("Tambah Nilai Peserta")
	fmt.Print("Masukkan ID : ")
	id := 0
	fmt.Scanln(&id)
	nilai1, nilai2, nilai3, nilai4 := 0.0,0.0,0.0,0.0
	fmt.Print("Masukkan Nilai 1 : ")
	fmt.Scanln(&nilai1)
	fmt.Print("Masukkan Nilai 2 : ")
	fmt.Scanln(&nilai2)
	fmt.Print("Masukkan Nilai 3 : ")
	fmt.Scanln(&nilai3)
	fmt.Print("Masukkan Nilai 4 : ")
	fmt.Scanln(&nilai4)
	PMB.TambahNilaiPeserta(id,nilai1,nilai2,nilai3,nilai4)
}

func searchParticipantByName(){
	fmt.Println("Cari Peserta")
	fmt.Print("Search : ")
	search := ""
	fmt.Scanln(&search)
	PMB.SearchName(search)
}

