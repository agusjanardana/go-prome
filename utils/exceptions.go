package utils

func PanicIfError(err error) {
	//defer EndApp()
	if err != nil {
		panic(err)
	}
	//fmt.Println("Aplikasi Berjalan")
}
