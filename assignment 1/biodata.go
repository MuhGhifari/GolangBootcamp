package main
  
import (
    "fmt"
    "os"
)

type Students struct {
    name    string
    address string
    job     string
    reason  string
}
  
func main() {
    
    // baca input dari user
    input := os.Args[1]
    var inputNumber int
    _, e:= fmt.Sscan(input, &inputNumber)

    // data murid di kelas
    class  := []Students{
        {"Dedi Chandra", "Yogyakarta", "Backend Developer", "Unknown"},
        {"Faikar Achmad Luthfi", "Tangerang", "Backend Engineer", "God's will"},
        {"Alvin Immanuel Simbolon", "Sumatera", "Backend Programmer", "Job requirement"},
        {"Muhammad Ghifari", "Bogor", "Backend Developer", "M.O.N.E.Y"},
        {"Khairul Abdi Dongoran", "Semarang", "Backend Engineer", "Loves learning"},
        {"Evrin Lumbantobing", "Jakarta", "Backend Programmer", "Obligation"},
    }
    
    if e != nil { // jika input tidak berbentuk integer
        fmt.Println("input harus berbentuk angka!")
    } else if len(class) == 0{ // jika data kelas kosong
        fmt.Println("Data kelas kosong!")
    } else {
        absen(class, inputNumber) // memanggil fungsi absen
    }
}

func absen(student []Students, absen int){
    
    // kondisi validasi
    isValidMin := absen > 0
    isValidMax := absen <= len(student)

    if isValidMin && isValidMax {
        index := absen - 1 // menyesuaikan input dengan index array
        fmt.Println("Nama        : ", student[index].name)
        fmt.Println("Alamat      : ", student[index].address)
        fmt.Println("Pekerjaan   : ", student[index].job)
        fmt.Println("Alasan      : ", student[index].reason)
    } else {
        if !isValidMin {
            fmt.Println("absen dimulai dari angka 1!")
        } 
        if !isValidMax {
            fmt.Println("angka melebihi jumlah murid di kelas!")
        }
    }

}