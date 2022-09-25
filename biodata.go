package main

import (
	"Assignment_1/src"
	"fmt"
	"os"
	"strconv"
)

var key_input *string = &os.Args[1]

func main() {
	absen, err := strconv.Atoi(*key_input)
	if err != nil {
		fmt.Printf("Tolong berikan input yang valid")
	} else {
		fmt.Printf("Berikut data siswa absen %d\n", absen)
		fmt.Print(src.Cek_absen(&absen))
	}
}