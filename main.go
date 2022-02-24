package main

import (
	"bufio" //dosya okurken kullanılıyor
	"fmt"
	"os"
	"strings"

	"github.com/yusufbu1ut/KitaplikApp/helper"
)

func main() {

	file, err := os.OpenFile("kitaplar.txt", os.O_RDONLY, 0755) //Dosya okunmak üzere açılır
	if err != nil {                                             //dizindeki dosyanın bulunamaması durumunda panic hata çıktısı verilir
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var bookSlice []string
	for scanner.Scan() { //burada gerçekleştirilen okuma değerleri slice içerisine aktarılır
		line := scanner.Text()
		bookSlice = append(bookSlice, line) //kitaplar slice içerisine yerleştirilir burada girdi sayısının belirsiz olduğu var sayılmıştır
	}

	if len(os.Args) < 2 {
		fmt.Println("Expected 'search' or 'list' subcommands")
		os.Exit(1)
	}
	//fmt.Println(os.Args)
	switch os.Args[1] {
	case "list":
		helper.List(bookSlice)
	case "search":
		srch := strings.Join(os.Args[2:], " ") //Search argumanları birleştirilir
		if len(os.Args) > 2 {
			helper.Search(srch, bookSlice)
		} else {
			fmt.Println("Expected argument for 'search'")
		}
	default:
		fmt.Println("Expected 'list' or 'serach' subcommands")
		os.Exit(1)
	}

	println("")
}
