package calculation

//untuk bisa diakses oleh package lain func harus diawali dengan huruf kapital
//jika hanya dipanggil dipackage yang sama dapat diawali huruf kecil
func Add(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}